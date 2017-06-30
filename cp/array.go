package cp

import (
	"github.com/v2pro/plz/acc"
)

func copierOfArray(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
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

func (copier *arrayCopier) Copy(dst interface{}, src interface{}) error {
	copier.srcAcc.IterateArray(src, func(elem interface{}) bool {
		copier.dstAcc.AppendArray(dst, func(dstElem interface{}) {
			copier.elemCopier.Copy(dstElem, elem)
		})
		return true
	})
	return nil
}
