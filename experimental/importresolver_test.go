package experimental_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/topxeq/gowasm"
	"github.com/topxeq/gowasm/api"
	"github.com/topxeq/gowasm/experimental"
	"github.com/topxeq/gowasm/internal/testing/binaryencoding"
	"github.com/topxeq/gowasm/internal/testing/require"
	"github.com/topxeq/gowasm/internal/wasm"
)

func TestImportResolver(t *testing.T) {
	ctx := context.Background()

	r := gowasm.NewRuntime(ctx)
	defer r.Close(ctx)

	for i := 0; i < 5; i++ {
		var callCount int
		start := func(ctx context.Context) {
			callCount++
		}
		modImport, err := r.NewHostModuleBuilder(fmt.Sprintf("env%d", i)).
			NewFunctionBuilder().WithFunc(start).Export("start").
			Compile(ctx)
		require.NoError(t, err)
		// Anonymous module, it will be resolved by the import resolver.
		instanceImport, err := r.InstantiateModule(ctx, modImport, gowasm.NewModuleConfig().WithName(""))
		require.NoError(t, err)

		resolveImport := func(name string) api.Module {
			if name == "env" {
				return instanceImport
			}
			return nil
		}

		// Set the import resolver in the context.
		ctx = experimental.WithImportResolver(context.Background(), resolveImport)

		one := uint32(1)
		binary := binaryencoding.EncodeModule(&wasm.Module{
			TypeSection:     []wasm.FunctionType{{}},
			ImportSection:   []wasm.Import{{Module: "env", Name: "start", Type: wasm.ExternTypeFunc, DescFunc: 0}},
			FunctionSection: []wasm.Index{0},
			CodeSection: []wasm.Code{
				{Body: []byte{wasm.OpcodeCall, 0, wasm.OpcodeEnd}}, // Call the imported env.start.
			},
			StartSection: &one,
		})

		modMain, err := r.CompileModule(ctx, binary)
		require.NoError(t, err)

		_, err = r.InstantiateModule(ctx, modMain, gowasm.NewModuleConfig())
		require.NoError(t, err)
		require.Equal(t, 1, callCount)
	}
}
