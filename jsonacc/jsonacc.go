package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
)

func init() {
	acc.Providers = append(acc.Providers, func(dstType reflect.Type, srcType reflect.Type) acc.Accessor {
		if reflect.TypeOf((*jsoniter.Iterator)(nil)) != srcType {
			return nil
		}
		if dstType.Kind() == reflect.Ptr {
			dstType = dstType.Elem()
		}
		switch dstType.Kind() {
		case reflect.Int:
			return &intAccessor{}
		case reflect.Map:
			elemAccessor := plz.AccessorOf(dstType.Elem(), srcType)
			return &mapAccessor{elemAccessor: elemAccessor}
		case reflect.Struct:
			elemAccessor := &emptyInterfaceAccessor{}
			return &mapAccessor{elemAccessor: elemAccessor}
		}
		return nil
	})
}
