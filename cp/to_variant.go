package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
)

type toVariantCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *toVariantCopier) Copy(dst, src unsafe.Pointer) error {
	dstElem, dstElemAcc := copier.dstAcc.InitVariant(dst, copier.srcAcc)
	elemCopier, err := util.CopierOf(dstElemAcc, copier.srcAcc)
	if err != nil {
		return err
	}
	return elemCopier.Copy(dstElem, src)
}
