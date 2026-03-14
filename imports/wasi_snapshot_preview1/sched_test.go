package wasi_snapshot_preview1_test

import (
	"testing"

	"github.com/topxeq/gowasm"
	"github.com/topxeq/gowasm/internal/testing/require"
	"github.com/topxeq/gowasm/internal/wasip1"
)

func Test_schedYield(t *testing.T) {
	var yielded bool
	mod, r, log := requireProxyModule(t, gowasm.NewModuleConfig().
		WithOsyield(func() {
			yielded = true
		}))
	defer r.Close(testCtx)
	requireErrnoResult(t, wasip1.ErrnoSuccess, mod, wasip1.SchedYieldName)
	require.Equal(t, `
==> wasi_snapshot_preview1.sched_yield()
<== errno=ESUCCESS
`, "\n"+log.String())
	require.True(t, yielded)
}
