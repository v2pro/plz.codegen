package acc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type int8Accessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *int8Accessor) Kind() lang.Kind {
	return lang.Int8
}

func (accessor *int8Accessor) GoString() string {
	return accessor.typ.Name()
}

func (accessor *int8Accessor) Int8(ptr unsafe.Pointer) int8 {
	return *(*int8)(ptr)
}

type ptrInt8Accessor struct {
	ptrAccessor
}

func (accessor *ptrInt8Accessor) Int8(ptr unsafe.Pointer) int8 {
	return accessor.valueAccessor.Int8(ptr)
}

func (accessor *ptrInt8Accessor) SetInt8(ptr unsafe.Pointer, val int8) {
	*((*int8)(ptr)) = val
}
