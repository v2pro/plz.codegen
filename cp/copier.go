package cp

import (
	"fmt"
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
)

func CopierOf(dstAcc lang.Accessor, srcAcc lang.Accessor) (util.Copier, error) {
	//srcKind := srcAcc.Kind()
	//dstKind := dstAcc.Kind()
	//if srcKind == lang.Map && dstKind == lang.Struct {
	//	return mapToStruct(dstAcc, srcAcc)
	//}
	//if srcKind == lang.Struct && dstKind == lang.Map {
	//	return structToMap(dstAcc, srcAcc)
	//}
	//if srcKind == lang.Variant {
	//	switch dstKind {
	//	case lang.String:
	//		return &stringCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//	case lang.Int:
	//		return &intCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//	}
	//	return &srcInterfaceCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//}
	//if dstKind == lang.Variant {
	//	switch srcKind {
	//	case lang.String:
	//		return &stringCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//	case lang.Int:
	//		return &intCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//	}
	//	return &dstInterfaceCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//}
	//if dstKind != srcKind {
	//	return nil, fmt.Errorf("kind mismatch: %#v => %#v", srcKind, dstKind)
	//}
	//switch srcKind {
	//case lang.Struct:
	//	return structToStruct(dstAcc, srcAcc)
	//case lang.String:
	//	return &stringCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//case lang.Int:
	//	return &intCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	//case lang.Map:
	//	return mapToMap(dstAcc, srcAcc)
	//case lang.Array:
	//	return arrayToArray(dstAcc, srcAcc)
	//}
	return nil, fmt.Errorf("do not know how to copy %#v => %#v", srcAcc, dstAcc)

}

//type stringCopier struct {
//	srcAcc lang.Accessor
//	dstAcc lang.Accessor
//}
//
//func (copier stringCopier) Copy(dst interface{}, src interface{}) error {
//	copier.dstAcc.SetString(dst, copier.srcAcc.String(src))
//	return nil
//}
//
//type intCopier struct {
//	srcAcc lang.Accessor
//	dstAcc lang.Accessor
//}
//
//func (copier intCopier) Copy(dst interface{}, src interface{}) error {
//	copier.dstAcc.SetInt(dst, copier.srcAcc.Int(src))
//	return nil
//}
//
//type srcInterfaceCopier struct {
//	srcAcc lang.Accessor
//	dstAcc lang.Accessor
//}
//
//func (copier srcInterfaceCopier) Copy(dst interface{}, src interface{}) error {
//	realSrc, realSrcAcc := copier.srcAcc.VariantElem(src)
//	if realSrc == nil {
//		return nil
//	}
//	realCopier, err := CopierOf(copier.dstAcc, realSrcAcc)
//	if err != nil {
//		return err
//	}
//	return realCopier.Copy(dst, realSrc)
//}
//
//type dstInterfaceCopier struct {
//	srcAcc lang.Accessor
//	dstAcc lang.Accessor
//}
//
//func (copier dstInterfaceCopier) Copy(dst interface{}, src interface{}) error {
//	realDst, realDstAcc := copier.dstAcc.VariantElem(dst)
//	if realDst == nil {
//		fmt.Println(reflect.TypeOf(dst))
//		fmt.Println(reflect.TypeOf(src))
//		panic("!!!")
//		realDst, realDstAcc = copier.dstAcc.InitVariant(dst, src)
//	}
//	realCopier, err := CopierOf(realDstAcc, copier.srcAcc)
//	if err != nil {
//		return err
//	}
//	return realCopier.Copy(realDst, src)
//}
