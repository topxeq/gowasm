package frontend

import (
	"github.com/topxeq/gowasm/internal/engine/wazevo/ssa"
	"github.com/topxeq/gowasm/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
