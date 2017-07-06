package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type stringCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *stringCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	val := copier.srcAcc.String(src)
	copier.dstAcc.SetString(dst, val)
	return nil
}
