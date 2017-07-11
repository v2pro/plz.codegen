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
func Compare_{{T}}(obj1 interface{}, obj2 interface{}) int {
	return typed_Compare_{{T}}(obj1.({{T}}), obj2.({{T}}))
}
func typed_Compare_{{T}}(obj1 {{T}}, obj2 {{T}}) int {
	if (obj1 < obj2) {
		return -1
	} else if (obj1 == obj2) {
		return 0
	} else {
		return 1
	}
}`,
		funcName: `Compare_{{T}}`,
	},
}

func Compare(obj1 interface{}, obj2 interface{}) int {
	typ := reflect.TypeOf(obj1)
	compare := compareSymbols.cache[typ]
	if compare == nil {
		typeName := typ.String()
		funcName, source := render(compareSymbols.template, `T`, typeName)
		compareObj := compile(source, funcName)
		compare = compareObj.(func(interface{}, interface{}) int)
		compareSymbols.cache[typ] = compare
	}
	return compare(obj1, obj2)
}
