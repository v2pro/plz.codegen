package cp
//
//import (
//	"github.com/v2pro/plz/lang"
//	"github.com/v2pro/plz/util"
//)
//
//func arrayToArray(dstAcc lang.Accessor, srcAcc lang.Accessor) (util.Copier, error) {
//	elemCopier, err := CopierOf(dstAcc.Elem(), srcAcc.Elem())
//	if err != nil {
//		return nil, err
//	}
//	return &arrayCopier{
//		elemCopier: elemCopier,
//		dstAcc:     dstAcc,
//		srcAcc:     srcAcc,
//	}, nil
//}
//
//type arrayCopier struct {
//	elemCopier util.Copier
//	dstAcc     lang.Accessor
//	srcAcc     lang.Accessor
//}
//
//func (copier *arrayCopier) Copy(dst interface{}, src interface{}) (err error) {
//	copier.dstAcc.FillArray(dst, func(filler lang.ArrayFiller) {
//		copier.srcAcc.IterateArray(src, func(index int, srcElem interface{}) bool {
//			_, dstElem := filler.Next()
//			if dstElem == nil {
//				return false
//			}
//			err = copier.elemCopier.Copy(dstElem, srcElem)
//			if err != nil {
//				return false
//			}
//			filler.Fill()
//			return true
//		})
//	})
//	return
//}
