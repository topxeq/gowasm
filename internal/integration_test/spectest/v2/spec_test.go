package v2

import (
	"context"
	"testing"

	"github.com/topxeq/gowasm"
	"github.com/topxeq/gowasm/api"
	"github.com/topxeq/gowasm/internal/integration_test/spectest"
	"github.com/topxeq/gowasm/internal/platform"
)

const enabledFeatures = api.CoreFeaturesV2

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), gowasm.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), gowasm.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}
