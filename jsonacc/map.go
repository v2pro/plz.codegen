package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
)

type mapAccessor struct {
	acc.NoopAccessor
	elemAccessor acc.Accessor
}

func (accessor *mapAccessor) Kind() reflect.Kind {
	return reflect.Map
}

func (accessor *mapAccessor) GoString() string {
	return "map"
}

func (accessor *mapAccessor) Key() acc.Accessor {
	return &mapKeyAccessor{}
}

func (accessor *mapAccessor) Elem() acc.Accessor {
	return accessor.elemAccessor
}

func (accessor *mapAccessor) IterateMap(obj interface{}, cb func(key interface{}, elem interface{}) bool) {
	iter := obj.(*jsoniter.Iterator)
	iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
		return cb(field, iter)
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
