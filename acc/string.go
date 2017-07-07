package acc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type stringAccessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *stringAccessor) Kind() lang.Kind {
	return lang.String
}

func (accessor *stringAccessor) GoString() string {
	return accessor.typ.Name()
}

func (accessor *stringAccessor) String(ptr unsafe.Pointer) string {
	return *((*string)(ptr))
}

type ptrStringAccessor struct {
	ptrAccessor
}

func (accessor *ptrStringAccessor) New() (interface{}, lang.Accessor) {
	obj := new(string)
	return obj, accessor
}

func (accessor *ptrStringAccessor) String(ptr unsafe.Pointer) string {
	return accessor.valueAccessor.String(ptr)
}

func (accessor *ptrStringAccessor) SetString(ptr unsafe.Pointer, val string) {
	*((*string)(ptr)) = val
}
