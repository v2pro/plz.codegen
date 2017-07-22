package compare

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

var ByItself = generic.DefineFunc("CompareByItself(val1 T, val2 T) int").
	Param("T", "the type of value to compare").
	Generators("dispatch", dispatch).
	Source(`
{{ $compare := expand (.T|dispatch) "T" .T }}
return {{$compare}}(val1, val2)`)

func dispatch(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "CompareSimpleValue"
	case reflect.Ptr:
		return "ComparePtr"
	}
	panic("unsupported type: " + typ.String())
}
