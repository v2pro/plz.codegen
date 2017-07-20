package compare

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
	"github.com/v2pro/plz"
)

var logger = plz.LoggerOf("package", "compare")

var ByField = generic.Func("CompareByField(val1 T, val2 T) int").
	Params(
	"T", "the type of value to compare",
	"F", "the field name").
	ImportFunc(ByItself).
	Generators("fieldOf", genFieldOf).
	Source(`
{{ $field := fieldOf .T .F }}
{{ $compare := expand "CompareByItself" "T" $field.Type }}
return {{$compare}}(val1.{{.F}}, val2.{{.F}})`)

func genFieldOf(typ reflect.Type, fieldName string) reflect.StructField {
	field, found := typ.FieldByName(fieldName)
	if !found {
		msg := "field " + fieldName + " not found in " + typ.String()
		logger.Error(nil, msg, "fieldName", fieldName, "type", typ.String())
		panic(msg)
	}
	return field
}
