package fp_max

import (
	"github.com/v2pro/plz/util"
	"reflect"
	_ "github.com/v2pro/wombat/acc"
	"unsafe"
	"fmt"
)

func init() {
	util.Max = max
}

func max(collection ...interface{}) interface{} {
	if len(collection) == 0 {
		return nil
	}
	typ := reflect.TypeOf(collection[0])
	f := funcMaxMap[typ]
	if f == nil {
		lastElem := collection[len(collection)-1]
		f = tryMaxStruct(typ, lastElem)
		if f != nil {
			return f.max(collection)
		}
		panic(fmt.Sprintf("no max implementation for: %v", typ))
	}
	return f.max(collection)
}

func objKind(obj interface{}) reflect.Kind {
	return reflect.Kind((*((*emptyInterface)(unsafe.Pointer(&obj)))).typ.kind & kindMask)
}

func objPtr(obj interface{}) unsafe.Pointer {
	return (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
}

const kindMask        = (1 << 5) - 1
type tflag uint8
// rtype is the common implementation of most values.
// It is embedded in other, public struct types, but always
// with a unique tag like `reflect:"array"` or `reflect:"ptr"`
// so that code cannot convert from, say, *arrayType to *ptrType.
type rtype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32   // hash of type; avoids computation in hash tables
	tflag      tflag    // extra type information flags
	align      uint8    // alignment of variable with this type
	fieldAlign uint8    // alignment of struct field with this type
	kind       uint8    // enumeration for C
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

var funcMaxMap = map[reflect.Type]funcMax{
	reflect.TypeOf(0): &maxInt{},
}

type funcMax interface {
	max(collection []interface{}) interface{}
}
