package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type skipCopier struct {
	srcAcc lang.Accessor
}

func (copier *skipCopier) Copy(dst, src unsafe.Pointer) error {
	copier.srcAcc.Skip(src)
	return nil
}