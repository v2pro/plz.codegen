package cpStructToStruct

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpStructToStruct"] = F
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cpStatically": cpStatically.F,
	},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $bindings := calcBindings (.DT|elem) .ST }}
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
	{{ range $_, $binding := $bindings }}
		{{ $binding.cp }}(err, &dst.{{ $binding.dstFieldName }}, src.{{ $binding.srcFieldName }})
	{{ end }}
}`,
	GenMap: map[string]interface{}{
		"calcBindings": genCalcBindings,
		"assignCp":     genAssignCp,
	},
}

func genCalcBindings(dstType, srcType reflect.Type) interface{} {
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

func genAssignCp(binding map[string]interface{}, cpFuncName string) string {
	binding["cp"] = cpFuncName
	return ""
}
