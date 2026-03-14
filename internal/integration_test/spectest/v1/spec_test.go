package v1

import (
	"context"
	"testing"

	"github.com/topxeq/gowasm"
	"github.com/topxeq/gowasm/api"
	"github.com/topxeq/gowasm/internal/integration_test/spectest"
	"github.com/topxeq/gowasm/internal/platform"
)

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), gowasm.NewRuntimeConfigCompiler().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), gowasm.NewRuntimeConfigInterpreter().WithCoreFeatures(api.CoreFeaturesV1))
}
