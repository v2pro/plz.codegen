package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/v2pro/plz/util"
)

type arrayCopier struct {
	srcAcc     lang.Accessor
	dstAcc     lang.Accessor
	elemCopier util.Copier
}

func (copier *arrayCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) (err error) {
	if src == nil {
		copier.dstAcc.Skip(dst)
		return nil
	}
	copier.dstAcc.FillArray(dst, func(filler lang.ArrayFiller) {
		copier.srcAcc.IterateArray(src, func(index int, srcElem unsafe.Pointer) bool {
			_, dstElem := filler.Next()
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
