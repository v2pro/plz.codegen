package acc

import (
	"fmt"
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type arrayAccessor struct {
	lang.NoopAccessor
	elemAcc lang.Accessor
	typ     reflect.Type
}

func (accessor *arrayAccessor) Kind() lang.Kind {
	return lang.Array
}

func (accessor *arrayAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *arrayAccessor) RandomAccessible() bool {
	return true
}

func (accessor *arrayAccessor) Elem() lang.Accessor {
	return accessor.elemAcc
}

func (accessor *arrayAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	if index < 0 {
		panic(fmt.Sprintf("index %v is negative", index))
	}
	if index >= accessor.typ.Len() {
		panic(fmt.Sprintf("index %v exceeded length %v", index, accessor.typ.Len()))
	}
	elemSize := accessor.typ.Elem().Size()
	elem := uintptr(ptr) + uintptr(index)*elemSize
	return unsafe.Pointer(elem)
}

func (accessor *arrayAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	elemSize := accessor.typ.Elem().Size()
	elemAcc := accessor.elemAcc
	head := uintptr(ptr)
	for index := 0; index < accessor.typ.Len(); index++ {
		elem := head + uintptr(index)*elemSize
		if elemAcc.IsNil(unsafe.Pointer(elem)) {
			elem = uintptr(0)
		}
		if !cb(index, unsafe.Pointer(elem)) {
			return
		}
	}
}

type ptrArrayAccessor struct {
	ptrAccessor
}

func (accessor *ptrArrayAccessor) Elem() lang.Accessor {
	typ := accessor.valueAccessor.(*arrayAccessor).typ
	return lang.AccessorOf(reflect.PtrTo(typ.Elem()))
}

func (accessor *ptrArrayAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	return accessor.valueAccessor.ArrayIndex(ptr, index)
}

func (accessor *ptrArrayAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	accessor.valueAccessor.IterateArray(ptr, cb)
}

func (accessor *ptrArrayAccessor) FillArray(ptr unsafe.Pointer, cb func(filler lang.ArrayFiller)) {
	typ := accessor.valueAccessor.(*arrayAccessor).typ
	filler := &arrayFiller{
		elemType: typ.Elem(),
		length:   typ.Len(),
		head:     uintptr(ptr),
	}
	cb(filler)
}

type arrayFiller struct {
	length   int
	elemType reflect.Type
	head     uintptr
	at       int
}

func (filler *arrayFiller) Next() (int, unsafe.Pointer) {
	at := filler.at
	if at >= filler.length {
		return -1, nil
	}
	elemPtr := filler.head + uintptr(at)*filler.elemType.Size()
	filler.at++
	return at, unsafe.Pointer(elemPtr)
}

func (filler *arrayFiller) Fill() {
}
