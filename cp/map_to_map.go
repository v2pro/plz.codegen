package cp

import "github.com/v2pro/plz/acc"

func mapToMap(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	keyCopier, err := CopierOf(dstAcc.Key(), srcAcc.Key())
	if err != nil {
		return nil, err
	}
	elemCopier, err := CopierOf(dstAcc.Elem(), srcAcc.Elem())
	if err != nil {
		return nil, err
	}
	return &mapCopier{
		keyCopier:  keyCopier,
		elemCopier: elemCopier,
		dstAcc:     dstAcc,
		srcAcc:     srcAcc,
	}, nil
}

type mapCopier struct {
	keyCopier  Copier
	elemCopier Copier
	dstAcc     acc.Accessor
	srcAcc     acc.Accessor
}

func (copier *mapCopier) Copy(dst interface{}, src interface{}) error {
	copier.dstAcc.FillMap(dst, func(filler acc.MapFiller) {
		copier.srcAcc.IterateMap(src, func(key interface{}, elem interface{}) bool {
			dstKey, dstElem := filler.Next()
			copier.keyCopier.Copy(dstKey, key)
			copier.elemCopier.Copy(dstElem, elem)
			filler.Fill()
			return true
		})
	})
	return nil
}
