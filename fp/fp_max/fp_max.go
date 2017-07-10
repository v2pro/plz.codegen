package fp_max

import (
	"github.com/v2pro/plz/util"
	"reflect"
	_ "github.com/v2pro/wombat/acc"
	"unsafe"
	"fmt"
	"github.com/v2pro/plz/lang"
)

func init() {
	util.Max = max
}

var intType = reflect.TypeOf(int(0))
var intMax funcMax = &funcMaxInt{}
var float64Type = reflect.TypeOf(float64(0))
var float64Max funcMax = &funcMaxFloat64{}
var int8Type = reflect.TypeOf(int8(0))
var int8Max funcMax = &funcMaxGeneric{comparator: lang.ObjectComparatorOf(reflect.Int8)}
var funcMaxMap = map[reflect.Type]funcMax{
	int8Type: int8Max,
}

func max(collection ...interface{}) interface{} {
	if len(collection) == 0 {
		return nil
	}
	typ := reflect.TypeOf(collection[0])
	if typ == intType {
		return intMax.max(collection)
	}
	if typ == float64Type {
		return float64Max.max(collection)
	}
	f := funcMaxMap[typ]
	if f != nil {
		return f.max(collection)
	}
	lastElem := collection[len(collection)-1]
	f = tryMaxStruct(typ, lastElem)
	if f != nil {
		return f.max(collection)
	}
	panic(fmt.Sprintf("no max implementation for: %v", typ))
	return f.max(collection)
}

func objKind(obj interface{}) reflect.Kind {
	return reflect.Kind((*((*emptyInterface)(unsafe.Pointer(&obj)))).typ.kind & kindMask)
}

func objPtr(obj interface{}) unsafe.Pointer {
	return (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
}

const kindMask = (1 << 5) - 1

// rtype is the common implementation of most values.
// It is embedded in other, public struct types, but always
// with a unique tag like `reflect:"array"` or `reflect:"ptr"`
// so that code cannot convert from, say, *arrayType to *ptrType.
type rtype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32 // hash of type; avoids computation in hash tables
	tflag      uint8  // extra type information flags
	align      uint8  // alignment of variable with this type
	fieldAlign uint8  // alignment of struct field with this type
	kind       uint8  // enumeration for C
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

type funcMax interface {
	max(collection []interface{}) interface{}
}
