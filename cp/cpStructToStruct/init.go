package cpStructToStruct

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpStructToStruct"] = F
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cpStatically": cpStatically.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $bindings := calcBindings (.DT|elem) .ST }}
{{ range $_, $binding := $bindings}}
	{{ $cp := gen "cpStatically" "DT" $binding.dstFieldType "ST" $binding.srcFieldType }}
	{{ $cp.Source }}
	{{ assignCp $binding $cp.FuncName }}
{{ end }}
func {{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	return typed_{{ .funcName }}(
		{{ cast "dst" .DT }},
		{{ cast "src" .ST }})
}
func typed_{{ .funcName }}(
	dst {{ .DT|name }},
	src {{ .ST|name }}) error {
	// end of signature
	{{ range $_, $binding := $bindings }}
		typed_{{ $binding.cp }}(&dst.{{ $binding.dstFieldName }}, src.{{ $binding.srcFieldName }})
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
