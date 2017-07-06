package cp
//
//import (
//	"github.com/v2pro/plz/lang"
//	"github.com/v2pro/plz/util"
//)
//
//func structToMap(dstAcc lang.Accessor, srcAcc lang.Accessor) (util.Copier, error) {
//	fieldCopiers, err := createStructToMapFieldCopiers(dstAcc, srcAcc)
//	if err != nil {
//		return nil, err
//	}
//	return &structToMapCopier{
//		fieldCopiers: fieldCopiers,
//		dstAcc:       dstAcc,
//		dstKeyAcc:    dstAcc.Key(),
//	}, nil
//}
//
//func createStructToMapFieldCopiers(dstAcc lang.Accessor, srcAcc lang.Accessor) (map[string]util.Copier, error) {
//	fieldCopiers := map[string]util.Copier{}
//	dstElemAcc := dstAcc.Elem()
//	for i := 0; i < srcAcc.NumField(); i++ {
//		field := srcAcc.Field(i)
//		copier, err := CopierOf(dstElemAcc, field.Accessor())
//		if err != nil {
//			return nil, err
//		}
//		fieldCopiers[field.Name()] = copier
//	}
//	return fieldCopiers, nil
//}
//
//type structToMapCopier struct {
//	fieldCopiers map[string]util.Copier
//	dstAcc       lang.Accessor
//	dstKeyAcc    lang.Accessor
//}
//
//func (copier *structToMapCopier) Copy(dst interface{}, src interface{}) (err error) {
//	copier.dstAcc.FillMap(dst, func(filler lang.MapFiller) {
//		for fieldName, fieldCopier := range copier.fieldCopiers {
//			dstKey, dstElem := filler.Next()
//			copier.dstKeyAcc.SetString(dstKey, fieldName)
//			err = fieldCopier.Copy(dstElem, src)
//			if err != nil {
//				return
//			}
//			filler.Fill()
//		}
//	})
//	return
//}
