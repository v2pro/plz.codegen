package cpMapToStruct

import (
	"github.com/v2pro/wombat/cp/cpAnything"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpAnything.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cpMapToStruct",
	Dependencies: []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $bindings := calcBindings (.DT|elem) .ST }}
{{ range $_, $binding := $bindings}}
	{{ $cp := gen "cpAnything" "DT" $binding.dstFieldType "ST" $binding.srcFieldType }}
	{{ assignCp $binding $cp }}
{{ end }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) error {
	// end of signature
	for key, elem := range src {
		switch key {
			{{ range $_, $binding := $bindings }}
				case "{{ $binding.srcFieldName }}":
					{{ $binding.cp }}(err, &dst.{{ $binding.dstFieldName }}, elem)
			{{ end }}
		}
	}
	return nil
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
		bindings = append(bindings, map[string]interface{}{
			"srcFieldName": dstField.Name,
			"srcFieldType": srcType.Elem(),
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
