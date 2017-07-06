package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
)

type fromVariantCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *fromVariantCopier) Copy(dst, src unsafe.Pointer) error {
	if src == nil {
		copier.dstAcc.Skip(dst)
		return nil
	}
	srcElem, srcElemAcc := copier.srcAcc.VariantElem(src)
	if srcElem == nil {
		copier.dstAcc.Skip(dst)
		return nil
	}
	elemCopier, err := util.CopierOf(copier.dstAcc, srcElemAcc)
	if err != nil {
		return err
	}
	return elemCopier.Copy(dst, srcElem)
}
