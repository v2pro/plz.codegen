package fp_compare

import (
	"reflect"
)

type structAndField struct {
	S reflect.Type
	F string
}

var compareStructByFieldSymbols = struct {
	template *funcTemplate
	cache    map[structAndField]func(interface{}, interface{}) int
}{
	cache: map[structAndField]func(interface{}, interface{}) int{},
	template: &funcTemplate{
		dependencies: []*funcTemplate{compareSymbols.template},
		variables: map[string]string{
			"S": "the struct type to compare",
			"F": "the field name of S",
			"T": "the type of field F",
		},
		source: `
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
	return typed_Compare_{{ .T|name }}(obj1.{{ .F }}, obj2.{{ .F }})
}`,
		funcName: `Compare_{{ .S|name }}_by_{{ .F }}`,
	},
}

func CompareStructByField(obj1 interface{}, obj2 interface{}, fieldName string) int {
	typ := reflect.TypeOf(obj1)
	cacheKey := structAndField{typ, fieldName}
	compare := compareStructByFieldSymbols.cache[cacheKey]
	if compare == nil {
		field, found := typ.FieldByName(fieldName)
		if !found {
			panic("field " + fieldName + " not found in " + typ.String())
		}
		funcName, source := gen(compareStructByFieldSymbols.template,
			`S`, typ, `F`, fieldName, `T`, field.Type)
		compareObj := compile(source, funcName)
		compare = compareObj.(func(interface{}, interface{}) int)
		compareStructByFieldSymbols.cache[cacheKey] = compare
	}
	return compare(obj1, obj2)
}
