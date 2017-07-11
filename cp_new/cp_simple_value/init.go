package cp_simple_value

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
)

var F = &gen.FuncTemplate{
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
func {{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) error {
	// end of signature
	return typed_{{ .funcName }}(
		obj1.({{ .DT|name }}),
		obj2.({{ .ST|name }}))
}
func typed_{{ .funcName }}(
	obj1 {{ .DT|name }},
	obj2 {{ .ST|name }}) error {
	// end of signature
	*obj1 = obj2
	return nil
}`,
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
