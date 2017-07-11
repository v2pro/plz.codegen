package compare

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
)

var F = &gen.FuncTemplate{
	Variables: map[string]string{
		"T": "the type to compare",
	},
	Source: `
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
	FuncName: `Compare_{{ .T|name }}`,
}

func Gen(typ reflect.Type) func(interface{}, interface{}) int {
	funcObj := gen.Compile(F, `T`, typ)
	return funcObj.(func(interface{}, interface{}) int)
}
