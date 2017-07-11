package cp_struct_to_struct

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp_new/cp_simple_value"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cp_simple_value": cp_simple_value.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $bindings := calcBindings (.DT|elem) .ST }}
{{ range $_, $binding := $bindings}}
	{{ $cp := gen "cp_simple_value" "DT" $binding.dstFieldType "ST" $binding.srcFieldType }}
	{{ $cp.Source }}
	{{ assignCp $binding $cp.FuncName }}
{{ end }}
func {{ .funcName }}(
	obj1 interface{},
	obj2 interface{}) error {
	// end of signature
	return typed_{{ .funcName }}(
		{{ cast "obj1" .DT }},
		{{ cast "obj2" .ST }})
}
func typed_{{ .funcName }}(
	obj1 {{ .DT|name }},
	obj2 *{{ .ST|name }}) error {
	// end of signature
	{{ range $_, $binding := $bindings }}
		typed_{{ $binding.cp }}(&obj1.{{ $binding.dstFieldName }}, obj2.{{ $binding.srcFieldName }})
	{{ end }}
	return nil
}`,
	FuncMap: map[string]interface{}{
		"calcBindings": calcBindings,
		"assignCp":     assignCp,
	},
}

func calcBindings(dstType, srcType reflect.Type) interface{} {
	bindings := []interface{}{}
	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		srcField, srcFieldFound := srcType.FieldByName(dstField.Name)
		if !srcFieldFound {
			continue
		}
		bindings = append(bindings, map[string]interface{}{
			"srcFieldName": srcField.Name,
			"srcFieldType": srcField.Type,
			"dstFieldName": dstField.Name,
			"dstFieldType": reflect.PtrTo(dstField.Type),
		})
	}
	return bindings
}

func assignCp(binding map[string]interface{}, cpFuncName string) string {
	binding["cp"] = cpFuncName
	return ""
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
