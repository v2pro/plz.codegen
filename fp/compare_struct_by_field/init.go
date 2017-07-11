package compare_struct_by_field

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/fp/compare"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"compare": compare.F,
	},
	Variables: map[string]string{
		"T": "the struct type to compare",
		"F": "the field name of T",
	},
	FuncName: `Compare_{{ .T|symbol }}_by_{{ .F }}`,
	Source: `
{{ $field := field_of .T .F }}
{{ $compareField := gen "compare" "T" $field.Type }}
{{ $compareField.Source }}
func {{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return typed_{{ .funcName }}({{ cast "obj1" .T }}, {{ cast "obj2" .T }})
}
func typed_{{ .funcName }}(
	obj1 *{{ .T|name }},
	obj2 *{{ .T|name }}) int {
	// end of signature
	return typed_{{ $compareField.FuncName }}(obj1.{{ .F }}, obj2.{{ .F }})
}`,
}

func Gen(typ reflect.Type, fieldName string) func(interface{}, interface{}) int {
	funcObj := gen.Compile(F, "T", typ, "F", fieldName)
	return funcObj.(func(interface{}, interface{}) int)
}
