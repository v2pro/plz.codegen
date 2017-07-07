package acc

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type ptrAccessor struct {
	lang.NoopAccessor
	valueAccessor lang.Accessor
}

func (accessor *ptrAccessor) Kind() lang.Kind {
	return accessor.valueAccessor.Kind()
}

func (accessor *ptrAccessor) GoString() string {
	return "*" + accessor.valueAccessor.GoString()
}

func (accessor *ptrAccessor) Key() lang.Accessor {
	return accessor.valueAccessor.Key()
}

func (accessor *ptrAccessor) Elem() lang.Accessor {
	return accessor.valueAccessor.Elem()
}

func (accessor *ptrAccessor) RandomAccessible() bool {
	return accessor.valueAccessor.RandomAccessible()
}

func (accessor *ptrAccessor) NumField() int {
	return accessor.valueAccessor.NumField()
}

func (accessor *ptrAccessor) Field(index int) lang.StructField {
	return accessor.valueAccessor.Field(index)
}

type ptrPtrAccessor struct {
	ptrAccessor
}

func (accessor *ptrPtrAccessor) New() (interface{}, lang.Accessor) {
	newObj, _ := accessor.valueAccessor.New()
	ptr := objPtr(newObj)
	return &ptr, accessor
}

func (accessor *ptrPtrAccessor) deRef(ptr unsafe.Pointer) unsafe.Pointer {
	realPtr := *((*unsafe.Pointer)(ptr))
	if realPtr == nil {
		newObj, _ := accessor.valueAccessor.New()
		realPtr = objPtr(newObj)
		*((*unsafe.Pointer)(ptr)) = realPtr
	}
	return realPtr
}

func (accessor *ptrPtrAccessor) IsNil(ptr unsafe.Pointer) bool {
	if ptr == nil {
		return true
	}
	return accessor.valueAccessor.IsNil(*((*unsafe.Pointer)(ptr)))
}

func (accessor *ptrPtrAccessor) Float64(ptr unsafe.Pointer) float64 {
	return accessor.valueAccessor.Float64(*((*unsafe.Pointer)(ptr)))
}

func (accessor *ptrPtrAccessor) SetFloat64(ptr unsafe.Pointer, val float64) {
	accessor.valueAccessor.SetFloat64(accessor.deRef(ptr), val)
}

func (accessor *ptrPtrAccessor) Int(ptr unsafe.Pointer) int {
	return accessor.valueAccessor.Int(*((*unsafe.Pointer)(ptr)))
}

func (accessor *ptrPtrAccessor) SetInt(ptr unsafe.Pointer, val int) {
	accessor.valueAccessor.SetInt(accessor.deRef(ptr), val)
}

func (accessor *ptrPtrAccessor) String(ptr unsafe.Pointer) string {
	return accessor.valueAccessor.String(*((*unsafe.Pointer)(ptr)))
}

func (accessor *ptrPtrAccessor) SetString(ptr unsafe.Pointer, val string) {
	accessor.valueAccessor.SetString(accessor.deRef(ptr), val)
}

func (accessor *ptrPtrAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	accessor.valueAccessor.IterateArray(*((*unsafe.Pointer)(ptr)), cb)
}

func (accessor *ptrPtrAccessor) FillArray(ptr unsafe.Pointer, cb func(filler lang.ArrayFiller)) {
	accessor.valueAccessor.FillArray(*((*unsafe.Pointer)(ptr)), cb)
}

func (accessor *ptrPtrAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	return accessor.valueAccessor.ArrayIndex(*((*unsafe.Pointer)(ptr)), index)
}

func (accessor *ptrPtrAccessor) IterateMap(ptr unsafe.Pointer, cb func(key unsafe.Pointer, elem unsafe.Pointer) bool) {
	accessor.valueAccessor.IterateMap(*((*unsafe.Pointer)(ptr)), cb)
}

func (accessor *ptrPtrAccessor) FillMap(ptr unsafe.Pointer, cb func(filler lang.MapFiller)) {
	accessor.valueAccessor.FillMap(accessor.deRef(ptr), cb)
}

func (accessor *ptrPtrAccessor) VariantElem(ptr unsafe.Pointer) (elem unsafe.Pointer, elemAccessor lang.Accessor) {
	realPtr := *((*unsafe.Pointer)(ptr))
	if realPtr == nil {
		return nil, nil
	}
	return accessor.valueAccessor.VariantElem(realPtr)
}

func (accessor *ptrPtrAccessor) InitVariant(ptr unsafe.Pointer, template lang.Accessor) (elem unsafe.Pointer, elemAccessor lang.Accessor) {
	return accessor.valueAccessor.InitVariant(accessor.deRef(ptr), template)
}
