package acc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type intAccessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *intAccessor) Kind() lang.Kind {
	return lang.Int
}

func (accessor *intAccessor) GoString() string {
	return accessor.typ.Name()
}

func (accessor *intAccessor) Int(ptr unsafe.Pointer) int {
	return *((*int)(ptr))
}
func (accessor *intAccessor) Compare(ptr1 unsafe.Pointer, ptr2 unsafe.Pointer) int {
	val1 := *(*int)(ptr1)
	val2 := *(*int)(ptr2)
	if val1 < val2 {
		return -1
	} else if val1 == val2 {
		return 0
	} else {
		return 1
	}
}

type ptrIntAccessor struct {
	ptrAccessor
}

func (accessor *ptrIntAccessor) New() (interface{}, lang.Accessor) {
	return new(int), accessor
}

func (accessor *ptrIntAccessor) Int(ptr unsafe.Pointer) int {
	return accessor.valueAccessor.Int(ptr)
}

func (accessor *ptrIntAccessor) SetInt(ptr unsafe.Pointer, val int) {
	*((*int)(ptr)) = val
}
