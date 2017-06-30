package cp

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"fmt"
)

func CopierOf(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	if dstAcc.Kind() == reflect.Struct && srcAcc.Kind() == reflect.Map {
		return mapToStruct(dstAcc, srcAcc)
	}
	if dstAcc.Kind() != srcAcc.Kind() {
		return nil, fmt.Errorf("kind mismatch: %v", dstAcc, srcAcc)
	}
	switch dstAcc.Kind() {
	case reflect.Struct:
		return structToStruct(dstAcc, srcAcc)
	case reflect.String:
		return &stringCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
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
