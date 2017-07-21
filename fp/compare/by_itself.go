package compare

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

var ByItself = generic.Func("CompareByItself(val1 T, val2 T) int").
	Param("T", "the type of value to compare").
	Generators("dispatch", dispatch).
	Source(`
{{ $compare := expand (.T|dispatch) "T" .T }}
return {{$compare}}(val1, val2)`)

func dispatch(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "CompareSimpleValue"
	case reflect.Ptr:
		return "ComparePtr"
	}
	panic("unsupported type: " + typ.String())
}
