package jsonacc

import (
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/lang"
)

type iteratorAccessor struct {
	lang.NoopAccessor
	kind lang.Kind
}

func (accessor *iteratorAccessor) Kind() lang.Kind {
	return accessor.kind
}

func (accessor *iteratorAccessor) GoString() string {
	return "iteratorAccessor"
}

func (accessor *iteratorAccessor) Key() lang.Accessor {
	return &mapKeyReader{}
}

func (accessor *iteratorAccessor) Elem() lang.Accessor {
	return &iteratorAccessor{kind: lang.Variant}
}

func (accessor *iteratorAccessor) PtrElem(obj interface{}) (interface{}, lang.Accessor) {
	iter := obj.(*jsoniter.Iterator)
	switch iter.WhatIsNext() {
	case jsoniter.Array:
		return obj, &iteratorAccessor{kind: lang.Array}
	case jsoniter.Object:
		return obj, &iteratorAccessor{kind: lang.Map}
	case jsoniter.Number:
		fallthrough
	case jsoniter.String:
		return obj, &iteratorAccessor{kind: lang.Variant}
	}
	panic("not implemented")
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

func (accessor *iteratorAccessor) IterateArray(obj interface{}, cb func(index int, elem interface{}) bool) {
	iter := obj.(*jsoniter.Iterator)
	index := 0
	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		currentIndex := index
		index++
		return cb(currentIndex, iter)
	})
}

type mapKeyReader struct {
	lang.NoopAccessor
}

func (accessor *mapKeyReader) Kind() lang.Kind {
	return lang.String
}

func (accessor *mapKeyReader) GoString() string {
	return "string"
}

func (accessor *mapKeyReader) String(obj interface{}) string {
	return obj.(string)
}
