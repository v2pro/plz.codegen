package cmpStructByField

import (
	"github.com/v2pro/wombat/fp/cmpSimpleValue"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cmpStructByField",
	Dependencies: []*gen.FuncTemplate{cmpSimpleValue.F},
	TemplateParams: map[string]string{
		"T": "the struct type to compare",
		"F": "the field name of T",
	},
	FuncName: `Compare_{{ .T|symbol }}_by_{{ .F }}`,
	Source: `
func Exported_{{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return {{ .funcName }}(
		{{ cast "obj1" .T }},
		{{ cast "obj2" .T }})
}
{{ if .T|isPtr }}
	{{ $compareElem := gen "cmpStructByField" "T" (.T|elem) "F" .F }}
	{{ $compareElem.Source }}
	func {{ .funcName }}(
		obj1 {{ .T|name }},
		obj2 {{ .T|name }}) int {
		// end of signature
		return {{ $compareElem.FuncName }}(*obj1, *obj2)
	}
{{ else }}
	{{ $field := fieldOf .T .F }}
	{{ $compareField := gen "cmpSimpleValue" "T" $field.Type }}
	{{ $compareField.Source }}
	func {{ .funcName }}(
		obj1 {{ .T|name }},
		obj2 {{ .T|name }}) int {
		// end of signature
		return {{ $compareField.FuncName }}(obj1.{{ .F }}, obj2.{{ .F }})
	}
{{ end }}`,
}

func genF(typ reflect.Type, fieldName string) func(interface{}, interface{}) int {
	funcObj := gen.Compile(F, "T", typ, "F", fieldName)
	return funcObj.(func(interface{}, interface{}) int)
}
