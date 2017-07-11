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
	Source: `
func {{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return typed_{{ .funcName }}({{ cast "obj1" .S }}, {{ cast "obj2" .S }})
}
func typed_{{ .funcName }}(
	obj1 {{ .S|name }},
	obj2 {{ .S|name }}) int {
	// end of signature
	return {{ .compareFuncName }}(obj1.{{ .F }}, obj2.{{ .F }})
}`,
	FuncName: `Compare_{{ .S|name }}_by_{{ .F }}`,
}

type structAndField struct {
	S reflect.Type
	F string
}

var symbols = map[structAndField]func(interface{}, interface{}) int{}

func DoCompareStructByField(obj1 interface{}, obj2 interface{}, fieldName string) int {
	typ := reflect.TypeOf(obj1)
	cacheKey := structAndField{typ, fieldName}
	f := symbols[cacheKey]
	if f == nil {
		field, found := typ.FieldByName(fieldName)
		if !found {
			panic("field " + fieldName + " not found in " + typ.String())
		}
		funcObj := gen.Compile(F,
			`S`, typ, `F`, fieldName, `T`, field.Type)
		f = funcObj.(func(interface{}, interface{}) int)
		symbols[cacheKey] = f
	}
	return f(obj1, obj2)
}
