package fp_compare

import (
	"reflect"
)

var compareSymbols = struct {
	template *funcTemplate
	cache    map[reflect.Type]func(interface{}, interface{}) int
}{
	cache: map[reflect.Type]func(interface{}, interface{}) int{},
	template: &funcTemplate{
		variables: map[string]string{
			"T": "the type to compare",
		},
		source: `
func {{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return typed_{{ .funcName }}(
		obj1.({{ .T|name }}),
		obj2.({{ .T|name }}))
}
func typed_{{ .funcName }}(
	obj1 {{ .T|name }},
	obj2 {{ .T|name }}) int {
	// end of signature
	if (obj1 < obj2) {
		return -1
	} else if (obj1 == obj2) {
		return 0
	} else {
		return 1
	}
}`,
		funcName: `Compare_{{ .T|name }}`,
	},
}

func Compare(obj1 interface{}, obj2 interface{}) int {
	typ := reflect.TypeOf(obj1)
	compare := compareSymbols.cache[typ]
	if compare == nil {
		funcName, source := gen(compareSymbols.template, `T`, typ)
		compareObj := compile(source, funcName)
		compare = compareObj.(func(interface{}, interface{}) int)
		compareSymbols.cache[typ] = compare
	}
	return compare(obj1, obj2)
}
