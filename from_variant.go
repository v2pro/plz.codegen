package wombat

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/v2pro/plz/util"
)

type fromVariantCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *fromVariantCopier) Copy(dst, src unsafe.Pointer) error {
	srcElem, srcElemAcc := copier.srcAcc.VariantElem(src)
	elemCopier, err := util.CopierOf(copier.dstAcc, srcElemAcc)
	if err != nil {
		return err
	}
	return elemCopier.Copy(dst, srcElem)
}
