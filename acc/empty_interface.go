package acc

import (
	"unsafe"
)

func castToEmptyInterface(obj interface{}) emptyInterface {
	return *((*emptyInterface)(unsafe.Pointer(&obj)))
}

func castBackEmptyInterface(ei emptyInterface) interface{} {
	return *((*interface{})(unsafe.Pointer(&ei)))
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
