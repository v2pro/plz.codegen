package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
)

func init() {
	acc.Providers = append(acc.Providers, func(typ reflect.Type, profile string) acc.Accessor {
		if "json" != profile {
			return nil
		}
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		switch typ.Kind() {
		case reflect.Int:
			return &intAccessor{}
		}
		return nil
	})
}

type intAccessor struct {
	acc.NoopAccessor
}

func (accessor *intAccessor) Kind() reflect.Kind {
	return reflect.Int
}

func (accessor *intAccessor) GoString() string {
	return "int"
}

func (accessor *intAccessor) Int(obj interface{}) int {
	iter := obj.(*jsoniter.Iterator)
	return iter.ReadInt()
}
