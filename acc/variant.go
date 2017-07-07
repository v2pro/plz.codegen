package acc

import (
	"fmt"
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type variantAccessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *variantAccessor) Kind() lang.Kind {
	return lang.Variant
}

func (accessor *variantAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *variantAccessor) VariantElem(ptr unsafe.Pointer) (unsafe.Pointer, lang.Accessor) {
	obj := *((*interface{})(ptr))
	if obj == nil {
		return nil, nil
	}
	typ := reflect.TypeOf(obj)
	return objPtr(obj), lang.AccessorOf(typ, accessor.TagName)
}

type ptrVariantAccessor struct {
	variantAccessor
}

func (accessor *ptrVariantAccessor) InitVariant(ptr unsafe.Pointer, template lang.Accessor) (elem unsafe.Pointer, elemAccessor lang.Accessor) {
	switch template.Kind() {
	case lang.String:
		return ptr, &stringVariantAccessor{
			lang.NoopAccessor{accessor.TagName, "stringVariantAccessor"},
			accessor.typ,
		}
	case lang.Int:
		return ptr, &intVariantAccessor{
			lang.NoopAccessor{accessor.TagName, "intVariantAccessor"},
			accessor.typ,
		}
	case lang.Float64:
		return ptr, &float64VariantAccessor{
			lang.NoopAccessor{accessor.TagName, "float64VariantAccessor"},
			accessor.typ,
		}
	case lang.Map:
		fallthrough
	case lang.Array:
		fallthrough
	case lang.Struct:
		obj, objAcc := template.New()
		*((*interface{})(ptr)) = obj
		return objPtr(obj), objAcc
	}
	panic(fmt.Sprintf("not implemented: %#v", template.Kind()))
}

type stringVariantAccessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *stringVariantAccessor) Kind() lang.Kind {
	return lang.String
}

func (accessor *stringVariantAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *stringVariantAccessor) SetString(ptr unsafe.Pointer, val string) {
	*((*interface{})(ptr)) = val
}

type intVariantAccessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *intVariantAccessor) Kind() lang.Kind {
	return lang.Int
}

func (accessor *intVariantAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *intVariantAccessor) SetInt(ptr unsafe.Pointer, val int) {
	*((*interface{})(ptr)) = val
}

type float64VariantAccessor struct {
	lang.NoopAccessor
	typ reflect.Type
}

func (accessor *float64VariantAccessor) Kind() lang.Kind {
	return lang.Float64
}

func (accessor *float64VariantAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *float64VariantAccessor) SetFloat64(ptr unsafe.Pointer, val float64) {
	*((*interface{})(ptr)) = val
}
