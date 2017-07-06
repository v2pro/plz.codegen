package cp
//
//import (
//	"github.com/v2pro/plz/lang"
//	"github.com/v2pro/plz/util"
//)
//
//func structToStruct(dstAcc lang.Accessor, srcAcc lang.Accessor) (util.Copier, error) {
//	fieldCopiers, err := createStructToStructFieldCopiers(dstAcc, srcAcc)
//	if err != nil {
//		return nil, err
//	}
//	return &structToStructCopier{fieldCopiers}, nil
//}
//
//func createStructToStructFieldCopiers(dstAcc lang.Accessor, srcAcc lang.Accessor) (map[string]Copier, error) {
//	bindings := map[string]*binding{}
//	for i := 0; i < dstAcc.NumField(); i++ {
//		field := dstAcc.Field(i)
//		bindings[field.Name()] = &binding{
//			dstAcc: field.Accessor(),
//		}
//	}
//	for i := 0; i < srcAcc.NumField(); i++ {
//		field := srcAcc.Field(i)
//		binding := bindings[field.Name()]
//		if binding == nil {
//			continue
//		}
//		binding.srcAcc = field.Accessor()
//	}
//	fieldCopiers := map[string]util.Copier{}
//	for fieldName, v := range bindings {
//		if v.srcAcc != nil && v.dstAcc != nil {
//			copier, err := CopierOf(v.dstAcc, v.srcAcc)
//			if err != nil {
//				return nil, err
//			}
//			fieldCopiers[fieldName] = copier
//		}
//	}
//	return fieldCopiers, nil
//}
//
//type binding struct {
//	srcAcc lang.Accessor
//	dstAcc lang.Accessor
//}
//
//type structToStructCopier struct {
//	fieldCopiers map[string]util.Copier
//}
//
//func (copier *structToStructCopier) Copy(dst interface{}, src interface{}) error {
//	for _, fieldCopier := range copier.fieldCopiers {
//		err := fieldCopier.Copy(dst, src)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
