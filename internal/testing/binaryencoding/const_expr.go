package binaryencoding

import (
	"github.com/topxeq/gowasm/internal/wasm"
)

func encodeConstantExpression(expr wasm.ConstantExpression) (ret []byte) {
	ret = append(ret, expr.Data...)
	return
}
