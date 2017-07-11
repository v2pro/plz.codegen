package max

import (
	"github.com/v2pro/wombat/gen"
	"reflect"
	"github.com/v2pro/wombat/fp/compare_struct_by_field"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"compareFuncName": compare_struct_by_field.F,
	},
	Variables: map[string]string{
		"S": "the struct type to compare",
		"F": "the field name of S",
		"T": "the type of field F",
	},
	FuncName: `Max_{{ .S|name }}_by_{{ .F }}`,
	Source: `
func {{ .funcName }}(objs []interface{}) interface{} {
	currentMaxObj := objs[0]
	for i := 1; i < len(objs); i++ {
		currentMax := {{ cast "currentMaxObj" .S }}
		elem := {{ cast "objs[i]" .S }}
		if typed_{{ .compareFuncName }}(elem, currentMax) > 0 {
			currentMaxObj = objs[i]
		}
	}
	return currentMaxObj
}
func typed_{{ .funcName }}(objs []{{ .S|name }}) {{ .S|name }} {
	currentMax := objs[0]
	for i := 1; i < len(objs); i++ {
		if {{ .compareFuncName }}(objs[i].{{ .F }}, currentMax.{{ .F }}) > 0 {
			currentMax = objs[i]
		}
	}
	return currentMax
}`,
}

type structAndField struct {
	S reflect.Type
	F string
}

var symbols = map[structAndField]func([]interface{}) interface{}{}

func Call(objs []interface{}, fieldName string) interface{} {
	typ := reflect.TypeOf(objs[0])
	cacheKey := structAndField{typ, fieldName}
	f := symbols[cacheKey]
	if f == nil {
		field, found := typ.FieldByName(fieldName)
		if !found {
			panic("field " + fieldName + " not found in " + typ.String())
		}
		funcObj := gen.Compile(F,
			`S`, typ, `F`, fieldName, `T`, field.Type)
		f = funcObj.(func([]interface{}) interface{})
		symbols[cacheKey] = f
	}
	return f(objs)
}
