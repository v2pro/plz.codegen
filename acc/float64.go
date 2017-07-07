package acc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type float64Accessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *float64Accessor) Kind() lang.Kind {
	return lang.Float64
}

func (accessor *float64Accessor) GoString() string {
	return accessor.typ.Name()
}

func (accessor *float64Accessor) Float64(ptr unsafe.Pointer) float64 {
	return *((*float64)(ptr))
}

type ptrFloat64Accessor struct {
	ptrAccessor
}

func (accessor *ptrFloat64Accessor) New() (interface{}, lang.Accessor) {
	return new(float64), accessor
}

func (accessor *ptrFloat64Accessor) Float64(ptr unsafe.Pointer) float64 {
	return accessor.valueAccessor.Float64(ptr)
}

func (accessor *ptrFloat64Accessor) SetFloat64(ptr unsafe.Pointer, val float64) {
	*((*float64)(ptr)) = val
}