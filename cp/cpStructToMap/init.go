package cpStructToMap

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpStructToMap"] = F
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cpStatically": cpStatically.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $bindings := calcBindings .DT .ST }}
{{ range $_, $binding := $bindings}}
	{{ $cp := gen "cpStatically" "DT" $binding.dstFieldType "ST" $binding.srcFieldType }}
	{{ $cp.Source }}
	{{ assignCp $binding $cp.FuncName }}
{{ end }}
// generated from cpStructToMap
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	var existingElem {{ .DT|elem|name }}
	var found bool
	{{ range $_, $binding := $bindings }}
		existingElem, found = dst["{{ $binding.dstFieldName }}"]
		if found {
			{{ $binding.cp }}(err, &existingElem, src.{{ $binding.srcFieldName }})
			dst["{{ $binding.dstFieldName }}"] = existingElem
		} else {
			newElem := new({{ $binding.dstFieldType|elem|name }})
			{{ $binding.cp }}(err, newElem, src.{{ $binding.srcFieldName }})
			dst["{{ $binding.dstFieldName }}"] = *newElem
		}
	{{ end }}
}`,
	FuncMap: map[string]interface{}{
		"calcBindings": calcBindings,
		"assignCp":     assignCp,
	},
}

func calcBindings(dstType, srcType reflect.Type) interface{} {
	bindings := []interface{}{}
	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		bindings = append(bindings, map[string]interface{}{
			"srcFieldName": srcField.Name,
			"srcFieldType": srcField.Type,
			"dstFieldName": srcField.Name,
			"dstFieldType": reflect.PtrTo(dstType.Elem()),
		})
	}
	return bindings
}

func assignCp(binding map[string]interface{}, cpFuncName string) string {
	binding["cp"] = cpFuncName
	return ""
}
