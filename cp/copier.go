package cp

import (
	"github.com/v2pro/plz/acc"
	"fmt"
)

func CopierOf(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	srcKind := srcAcc.Kind()
	dstKind := dstAcc.Kind()
	if (srcKind == acc.Map || srcKind == acc.Interface) && dstKind == acc.Struct {
		return mapToStruct(dstAcc, srcAcc)
	}
	if srcKind == acc.Struct && (dstKind == acc.Map || dstKind == acc.Interface) {
		return structToMap(dstAcc, srcAcc)
	}
	if srcKind != acc.Interface && dstKind != acc.Interface && dstKind != srcKind {
		return nil, fmt.Errorf("kind mismatch: %#v => %#v", srcKind, dstKind)
	}
	kind := srcKind
	if kind == acc.Interface {
		kind = dstKind
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
	srcAcc := copier.srcAcc.AccessorOf(src)
	realCopier, err := CopierOf(copier.dstAcc, srcAcc)
	if err != nil {
		return err
	}
	return realCopier.Copy(dst, src)
}
