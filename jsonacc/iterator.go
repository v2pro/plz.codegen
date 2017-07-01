package jsonacc

import (
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/acc"
)

type iteratorAccessor struct {
	acc.NoopAccessor
}

func (accessor *iteratorAccessor) Kind() acc.Kind {
	return acc.Interface
}

func (accessor *iteratorAccessor) GoString() string {
	return "interface{}"
}

func (accessor *iteratorAccessor) Key() acc.Accessor {
	return &mapKeyReader{}
}

func (accessor *iteratorAccessor) Elem() acc.Accessor {
	return accessor
}

func (accessor *iteratorAccessor) Int(obj interface{}) int {
	iter := obj.(*jsoniter.Iterator)
	return iter.ReadInt()
}

func (accessor *iteratorAccessor) String(obj interface{}) string {
	iter := obj.(*jsoniter.Iterator)
	return iter.ReadString()
}

func (accessor *iteratorAccessor) IterateMap(obj interface{}, cb func(key interface{}, elem interface{}) bool) {
	iter := obj.(*jsoniter.Iterator)
	iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
		return cb(field, iter)
	})
}

func (accessor *iteratorAccessor) IterateArray(obj interface{}, cb func(elem interface{}) bool) {
	iter := obj.(*jsoniter.Iterator)
	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		return cb(iter)
	})
}

type mapKeyReader struct {
	acc.NoopAccessor
}

func (accessor *mapKeyReader) Kind() acc.Kind {
	return acc.String
}

func (accessor *mapKeyReader) GoString() string {
	return "string"
}

func (accessor *mapKeyReader) String(obj interface{}) string {
	return obj.(string)
}

