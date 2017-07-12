package cmpSimpleValue

import (
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	F.Dependencies["cmpSimpleValue"] = F
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
	// set in init()
	// "cmpSimpleValue": F,
	},
	Variables: map[string]string{
		"T": "the type to compare",
	},
	FuncName: `Compare_{{ .T|symbol }}`,
	Source: `
{{ if .T|isPtr }}
	{{ $compareElem := gen "cmpSimpleValue" "T" (.T|elem) }}
	{{ $compareElem.Source }}
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
		return typed_{{ $compareElem.FuncName }}(*obj1, *obj2)
	}
{{ else }}
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
	}
{{ end }}
`,
}

func Gen(typ reflect.Type) func(interface{}, interface{}) int {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Int, reflect.Int8, reflect.Int16:
		funcObj := gen.Compile(F, `T`, typ)
		return funcObj.(func(interface{}, interface{}) int)
	}
	return nil
}
