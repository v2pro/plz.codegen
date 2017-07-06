package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type float64Copier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *float64Copier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	copier.dstAcc.SetFloat64(dst, copier.srcAcc.Float64(src))
	return nil
}
