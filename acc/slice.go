package acc

import (
	"fmt"
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type sliceAccessor struct {
	lang.NoopAccessor
	elemAcc lang.Accessor
	typ     reflect.Type
}

func (accessor *sliceAccessor) Kind() lang.Kind {
	return lang.Array
}

func (accessor *sliceAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *sliceAccessor) Elem() lang.Accessor {
	return accessor.elemAcc
}

func (accessor *sliceAccessor) RandomAccessible() bool {
	return true
}

func (accessor *sliceAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	sliceHeader := (*sliceHeader)(ptr)
	if index < 0 {
		panic(fmt.Sprintf("index %v is negative", index))
	}
	if index >= sliceHeader.Len {
		panic(fmt.Sprintf("index %v exceeded length %v", index, sliceHeader.Len))
	}
	elemSize := accessor.typ.Elem().Size()
	head := uintptr(sliceHeader.Data)
	elem := head + uintptr(index)*elemSize
	return unsafe.Pointer(elem)
}

func (accessor *sliceAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	sliceHeader := (*sliceHeader)(ptr)
	elemSize := accessor.typ.Elem().Size()
	head := uintptr(sliceHeader.Data)
	for index := 0; index < sliceHeader.Len; index++ {
		elem := head + uintptr(index)*elemSize
		if accessor.elemAcc.IsNil(unsafe.Pointer(elem)) {
			elem = uintptr(0)
		}
		if !cb(index, unsafe.Pointer(elem)) {
			return
		}
	}
}

func (accessor *sliceAccessor) New() (interface{}, lang.Accessor) {
	return reflect.New(accessor.typ).Elem().Interface(), lang.AccessorOf(reflect.PtrTo(accessor.typ))
}

type ptrSliceAccessor struct {
	ptrAccessor
}

func (accessor *ptrSliceAccessor) Elem() lang.Accessor {
	typ := accessor.valueAccessor.(*sliceAccessor).typ
	return lang.AccessorOf(reflect.PtrTo(typ.Elem()))
}

func (accessor *ptrSliceAccessor) New() (interface{}, lang.Accessor) {
	typ := accessor.valueAccessor.(*sliceAccessor).typ
	return reflect.New(typ).Elem().Interface(), accessor
}

func (accessor *ptrSliceAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	return accessor.valueAccessor.ArrayIndex(ptr, index)
}

func (accessor *ptrSliceAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	accessor.valueAccessor.IterateArray(ptr, cb)
}

func (accessor *ptrSliceAccessor) FillArray(ptr unsafe.Pointer, cb func(filler lang.ArrayFiller)) {
	typ := accessor.valueAccessor.(*sliceAccessor).typ
	header := *(*sliceHeader)(ptr)
	header.Len = 0
	filler := &sliceFiller{
		sliceType: typ,
		elemType:  typ.Elem(),
		header:    &header,
	}
	cb(filler)
	*(*sliceHeader)(ptr) = *filler.header
}

type sliceFiller struct {
	sliceType reflect.Type
	elemType  reflect.Type
	header    *sliceHeader
}

func (filler *sliceFiller) Next() (int, unsafe.Pointer) {
	header := filler.header
	at := header.Len
	growOne(header, filler.sliceType, filler.elemType)
	elemPtr := uintptr(header.Data) + uintptr(at)*filler.elemType.Size()
	return at, unsafe.Pointer(elemPtr)
}

func (filler *sliceFiller) Fill() {
}

// grow grows the slice s so that it can hold extra more values, allocating
// more capacity if needed. It also returns the old and new slice lengths.
func growOne(slice *sliceHeader, sliceType reflect.Type, elementType reflect.Type) {
	newLen := slice.Len + 1
	if newLen <= slice.Cap {
		slice.Len = newLen
		return
	}
	newCap := slice.Cap
	if newCap == 0 {
		newCap = 1
	} else {
		for newCap < newLen {
			if slice.Len < 1024 {
				newCap += newCap
			} else {
				newCap += newCap / 4
			}
		}
	}
	dst := unsafe.Pointer(reflect.MakeSlice(sliceType, newLen, newCap).Pointer())
	// copy old array into new array
	originalBytesCount := uintptr(slice.Len) * elementType.Size()
	srcPtr := (*[1 << 30]byte)(slice.Data)
	dstPtr := (*[1 << 30]byte)(dst)
	for i := uintptr(0); i < originalBytesCount; i++ {
		dstPtr[i] = srcPtr[i]
	}
	slice.Len = newLen
	slice.Cap = newCap
	slice.Data = dst
}

// sliceHeader is a safe version of SliceHeader used within this package.
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
