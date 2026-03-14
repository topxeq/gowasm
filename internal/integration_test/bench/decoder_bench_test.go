package bench

import (
	"testing"

	"github.com/topxeq/gowasm/api"
	"github.com/topxeq/gowasm/internal/wasm"
	"github.com/topxeq/gowasm/internal/wasm/binary"
)

func BenchmarkCodec(b *testing.B) {
	b.Run("binary.DecodeModule", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if _, err := binary.DecodeModule(caseWasm, api.CoreFeaturesV2, wasm.MemoryLimitPages, false, false, false); err != nil {
				b.Fatal(err)
			}
		}
	})
}
