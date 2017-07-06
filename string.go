package wombat

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type stringCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *stringCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	copier.dstAcc.SetString(dst, copier.srcAcc.String(src))
	return nil
}