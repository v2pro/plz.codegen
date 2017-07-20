package compare

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

var F = generic.Func("Compare").
	Params("T", "the type of value to compare").
	Generators("dispatch", dispatch).
	ImportFunc(compareSimpleValue).
	Source(`
{{ $compare := expand (.T|dispatch) "T" .T }}
func {{.funcName}}(val1 {{.T|name}}, val2 {{.T|name}}) int {
	return {{$compare}}(val1, val2)
}`)


func dispatch(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "CompareSimpleValue"
	case reflect.Ptr:
		return "ComparePtr"
	}
	panic("unsupported type: " + typ.String())
}