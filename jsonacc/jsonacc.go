package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
)

func init() {
	acc.Providers = append(acc.Providers, func(typ reflect.Type) acc.Accessor {
		if reflect.TypeOf((*jsoniter.Iterator)(nil)) != typ {
			return nil
		}
		return &iterAcc{}
	})
}

type iterAcc struct {
	acc.NoopAccessor
}

func (accessor *iterAcc) Kind() reflect.Kind {
	return reflect.Interface
}

func (accessor *iterAcc) GoString() string {
	return "interface{}"
}

func (accessor *iterAcc) Int(obj interface{}) int {
	iter := obj.(*jsoniter.Iterator)
	return iter.ReadInt()
}

func (accessor *iterAcc) Key() acc.Accessor {
	return &mapKeyAccessor{}
}

func (accessor *iterAcc) Elem() acc.Accessor {
	return accessor
}

func (accessor *iterAcc) IterateMap(obj interface{}, cb func(key interface{}, elem interface{}) bool) {
	iter := obj.(*jsoniter.Iterator)
	iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
		return cb(field, iter)
	})
}

func (accessor *iterAcc) IterateArray(obj interface{}, cb func(elem interface{}) bool) {
	iter := obj.(*jsoniter.Iterator)
	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		return cb(iter)
	})
}

type mapKeyAccessor struct {
	acc.NoopAccessor
}

func (accessor *mapKeyAccessor) Kind() reflect.Kind {
	return reflect.String
}

func (accessor *mapKeyAccessor) GoString() string {
	return "string"
}

func (accessor *mapKeyAccessor) String(obj interface{}) string {
	return obj.(string)
}
