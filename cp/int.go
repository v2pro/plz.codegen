package cp

import (
	"unsafe"
	"github.com/v2pro/plz/lang"
)

type intCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *intCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	copier.dstAcc.SetInt(dst, copier.srcAcc.Int(src))
	return nil
}