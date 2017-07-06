package cp
//
//import (
//	"github.com/v2pro/plz/lang"
//	"github.com/v2pro/plz/util"
//)
//
//func mapToStruct(dstAcc lang.Accessor, srcAcc lang.Accessor) (util.Copier, error) {
//	fieldCopiers, err := createMapToStructFieldCopiers(dstAcc, srcAcc)
//	if err != nil {
//		return nil, err
//	}
//	return &mapToStructCopier{
//		fieldCopiers: fieldCopiers,
//		srcAcc: srcAcc,
//		srcKeyAcc: srcAcc.Key(),
//		srcElemAcc: srcAcc.Elem(),
//	}, nil
//}
//
//func createMapToStructFieldCopiers(dstAcc lang.Accessor, srcAcc lang.Accessor) (map[string]util.Copier, error) {
//	fieldCopiers := map[string]util.Copier{}
//	for i := 0; i < dstAcc.NumField(); i++ {
//		field := dstAcc.Field(i)
//		var err error
//		fieldCopiers[field.Name()], err = CopierOf(field.Accessor(), srcAcc.Elem())
//		if err != nil {
//			return nil, err
//		}
//	}
//	return fieldCopiers, nil
//}
//
//type mapToStructCopier struct {
//	fieldCopiers map[string]util.Copier
//	srcAcc       lang.Accessor
//	srcKeyAcc    lang.Accessor
//	srcElemAcc   lang.Accessor
//}
//
//func (copier *mapToStructCopier) Copy(dst interface{}, src interface{}) error {
//	copier.srcAcc.IterateMap(src, func(key interface{}, elem interface{}) bool {
//		fieldName := copier.srcKeyAcc.String(key)
//		fieldCopier := copier.fieldCopiers[fieldName]
//		if fieldCopier == nil {
//			copier.srcElemAcc.Skip(elem)
//		} else {
//			fieldCopier.Copy(dst, elem)
//		}
//		return true
//	})
//	return nil
//}
