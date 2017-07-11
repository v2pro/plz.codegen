package compare_struct_by_field

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/fp/compare"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"compareFuncName": compare.F,
	},
	Variables: map[string]string{
		"S": "the struct type to compare",
		"F": "the field name of S",
		"T": "the type of field F",
	},
	FuncName: `Compare_{{ .S|name }}_by_{{ .F }}`,
	Source: `
func {{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return typed_{{ .funcName }}({{ cast "obj1" .S }}, {{ cast "obj2" .S }})
}
func typed_{{ .funcName }}(
	obj1 *{{ .S|name }},
	obj2 *{{ .S|name }}) int {
	// end of signature
	return typed_{{ .compareFuncName }}(obj1.{{ .F }}, obj2.{{ .F }})
}`,
}

func Gen(typ reflect.Type, fieldName string) func(interface{}, interface{}) int {
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic("field " + fieldName + " not found in " + typ.String())
	}
	funcObj := gen.Compile(F,
		`S`, typ, `F`, fieldName, `T`, field.Type)
	return funcObj.(func(interface{}, interface{}) int)
}
