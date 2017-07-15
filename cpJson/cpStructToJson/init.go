package cpStructToJson

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpStructToJson"] = F
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
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	dst.WriteObjectStart()
	{{ range $i, $binding := $bindings }}
		{{ if ne $i 0 }}
			dst.WriteMore()
		{{ end }}
		dst.WriteObjectField("{{ $binding.dstFieldName }}")
		{{ $binding.cp }}(err, dst, src.{{ $binding.srcFieldName }})
	{{ end }}
	dst.WriteObjectEnd()
}
`,
	GenMap: map[string]interface{}{
		"calcBindings": genCalcBindings,
		"assignCp":     genAssignCp,
	},
}

func genCalcBindings(dstType, srcType reflect.Type) interface{} {
	bindings := []interface{}{}
	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		bindings = append(bindings, map[string]interface{}{
			"srcFieldName": srcField.Name,
			"srcFieldType": srcField.Type,
			"dstFieldName": srcField.Name,
			"dstFieldType": dstType,
		})
	}
	return bindings
}

func genAssignCp(binding map[string]interface{}, cpFuncName string) string {
	binding["cp"] = cpFuncName
	return ""
}
