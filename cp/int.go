package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type intCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *intCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	if dst == nil {
		copier.srcAcc.Skip(src)
		return nil
	}
	if src == nil {
		copier.dstAcc.Skip(dst)
		return nil
	}
	copier.dstAcc.SetInt(dst, copier.srcAcc.Int(src))
	return nil
}
