package wombat

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
)

func newMapToMapCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
	keyCopier, err := util.CopierOf(dstAcc.Key(), srcAcc.Key())
	if err != nil {
		return nil, err
	}
	elemCopier, err := util.CopierOf(dstAcc.Elem(), srcAcc.Elem())
	if err != nil {
		return nil, err
	}
	return &mapToMapCopier{
		srcAcc:     srcAcc,
		dstAcc:     dstAcc,
		keyCopier:  keyCopier,
		elemCopier: elemCopier,
	}, nil
}

type mapToMapCopier struct {
	srcAcc     lang.Accessor
	dstAcc     lang.Accessor
	keyCopier  util.Copier
	elemCopier util.Copier
}

func (copier *mapToMapCopier) Copy(dst, src unsafe.Pointer) (err error) {
	copier.dstAcc.FillMap(dst, func(filler lang.MapFiller) {
		copier.srcAcc.IterateMap(src, func(srcKey unsafe.Pointer, srcElem unsafe.Pointer) bool {
			dstKey, dstElem := filler.Next()
			err = copier.keyCopier.Copy(dstKey, srcKey)
			if err != nil {
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
