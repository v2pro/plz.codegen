package cp

import (
	"github.com/v2pro/plz/acc"
	"fmt"
)

func CopierOf(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	if (srcAcc.Kind() == acc.Map || srcAcc.Kind() == acc.Interface) && dstAcc.Kind() == acc.Struct {
		return mapToStruct(dstAcc, srcAcc)
	}
	if srcAcc.Kind() == acc.Struct && (dstAcc.Kind() == acc.Map || dstAcc.Kind() == acc.Interface) {
		return structToMap(dstAcc, srcAcc)
	}
	if srcAcc.Kind() != acc.Interface && dstAcc.Kind() != acc.Interface && dstAcc.Kind() != srcAcc.Kind() {
		return nil, fmt.Errorf("kind mismatch: %#v => %#v", srcAcc.Kind(), dstAcc.Kind())
	}
	kind := srcAcc.Kind()
	if kind == acc.Interface {
		kind = dstAcc.Kind()
	}
	switch kind {
	case acc.Struct:
		return structToStruct(dstAcc, srcAcc)
	case acc.String:
		return &stringCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	case acc.Int:
		return &intCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	case acc.Interface:
		return &interfaceCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	case acc.Map:
		return mapToMap(dstAcc, srcAcc)
	case acc.Array:
		return arrayToArray(dstAcc, srcAcc)
	default:
		return nil, fmt.Errorf("do not know how to copy %#v => %#v", srcAcc, dstAcc)
	}
}

type Copier interface {
	Copy(dst interface{}, src interface{}) error
}

type stringCopier struct {
	srcAcc acc.Accessor
	dstAcc acc.Accessor
}

func (copier stringCopier) Copy(dst interface{}, src interface{}) error {
	copier.dstAcc.SetString(dst, copier.srcAcc.String(src))
	return nil
}

type intCopier struct {
	srcAcc acc.Accessor
	dstAcc acc.Accessor
}

func (copier intCopier) Copy(dst interface{}, src interface{}) error {
	copier.dstAcc.SetInt(dst, copier.srcAcc.Int(src))
	return nil
}

type interfaceCopier struct {
	srcAcc acc.Accessor
	dstAcc acc.Accessor
}

func (copier interfaceCopier) Copy(dst interface{}, src interface{}) error {
	kind := copier.srcAcc.KindOf(src)
	switch kind {
	case acc.Int:
		copier.dstAcc.SetInt(dst, copier.srcAcc.Int(src))
	case acc.String:
		copier.dstAcc.SetString(dst, copier.srcAcc.String(src))
	}
	return fmt.Errorf("do not know how to copy %v at runtime", kind)
}
