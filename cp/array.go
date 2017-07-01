package cp

import (
	"github.com/v2pro/plz/acc"
)

func arrayToArray(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	elemCopier, err := CopierOf(dstAcc.Elem(), srcAcc.Elem())
	if err != nil {
		return nil, err
	}
	return &arrayCopier{
		elemCopier: elemCopier,
		dstAcc:     dstAcc,
		srcAcc:     srcAcc,
	}, nil
}

type arrayCopier struct {
	elemCopier Copier
	dstAcc     acc.Accessor
	srcAcc     acc.Accessor
}

func (copier *arrayCopier) Copy(dst interface{}, src interface{}) (err error) {
	copier.dstAcc.FillArray(dst, func(filler acc.ArrayFiller) {
		copier.srcAcc.IterateArray(src, func(srcElem interface{}) bool {
			dstElem := filler.Next()
			if dstElem == nil {
				return false
			}
			err = copier.elemCopier.Copy(dstElem, srcElem)
			if err != nil {
				return false
			}
			filler.Fill()
			return true
		})
	})
	return
}
