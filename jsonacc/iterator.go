package jsonacc

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/json-iterator/go"
)

type iteratorAccessor struct {
	lang.NoopAccessor
	kind lang.Kind
}

func (accessor *iteratorAccessor) ReadOnly() bool {
	return true
}

func (accessor *iteratorAccessor) Kind() lang.Kind {
	return accessor.kind
}

func (accessor *iteratorAccessor) GoString() string {
	return "iteratorAccessor"
}

//
//func (accessor *iteratorAccessor) Key() lang.Accessor {
//	return &mapKeyReader{}
//}
//
//func (accessor *iteratorAccessor) Elem() lang.Accessor {
//	return &iteratorAccessor{kind: lang.Variant}
//}
//
func (accessor *iteratorAccessor) VariantElem(ptr unsafe.Pointer) (unsafe.Pointer, lang.Accessor) {
	iter := (*jsoniter.Iterator)(ptr)
	switch iter.WhatIsNext() {
	case jsoniter.Array:
		return ptr, &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			lang.Array,
		}
	case jsoniter.Object:
		return ptr, &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			lang.Map,
		}
	case jsoniter.Number:
		return ptr, &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			lang.Float64,
		}
	case jsoniter.String:
		return ptr, &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			lang.String,
		}
	}
	panic("not implemented")
}

func (accessor *iteratorAccessor) Int(ptr unsafe.Pointer) int {
	iter := (*jsoniter.Iterator)(ptr)
	return iter.ReadInt()
}

func (accessor *iteratorAccessor) Float64(ptr unsafe.Pointer) float64 {
	iter := (*jsoniter.Iterator)(ptr)
	return iter.ReadFloat64()
}

//
//func (accessor *iteratorAccessor) String(obj unsafe.Pointer) string {
//	iter := obj.(*jsoniter.Iterator)
//	return iter.ReadString()
//}
//
//func (accessor *iteratorAccessor) IterateMap(obj interface{}, cb func(key interface{}, elem interface{}) bool) {
//	iter := obj.(*jsoniter.Iterator)
//	iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
//		return cb(field, iter)
//	})
//}
//
//func (accessor *iteratorAccessor) IterateArray(obj interface{}, cb func(index int, elem interface{}) bool) {
//	iter := obj.(*jsoniter.Iterator)
//	index := 0
//	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
//		currentIndex := index
//		index++
//		return cb(currentIndex, iter)
//	})
//}
//
//type mapKeyReader struct {
//	lang.NoopAccessor
//}
//
//func (accessor *mapKeyReader) ReadOnly() bool {
//	return true
//}
//
//func (accessor *mapKeyReader) Kind() lang.Kind {
//	return lang.String
//}
//
//func (accessor *mapKeyReader) GoString() string {
//	return "string"
//}
//
//func (accessor *mapKeyReader) String(obj interface{}) string {
//	return obj.(string)
//}
