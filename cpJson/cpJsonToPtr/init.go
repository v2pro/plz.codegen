package cpJsonToPtr

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cpJsonToPtr",
	Dependencies: []*gen.FuncTemplate{cpStatically.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cpStatically" "DT" (.DT|elem) "ST" .ST }}
{{ $cp.Source }}
// generated from cpJsonToPtr
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	if dst == nil {
		src.Skip()
		return
	}
	if src.ReadNil() {
		*dst = nil
		return
	}
	defDst := *dst
	if defDst == nil {
		{{ if .DT|elem|isMap }}
			defDst = {{ .DT|elem|name }}{}
		{{ else }}
			defDst = new({{ .DT|elem|elem|name }})
		{{ end }}
		{{ $cp.FuncName }}(err, defDst, src)
		*dst = defDst
		return
	}
	{{ $cp.FuncName }}(err, *dst, src)
}
`,
	GenMap: map[string]interface{}{
		"isMap": genIsMap,
	},
}

func genIsMap(typ reflect.Type) bool {
	return typ.Kind() == reflect.Map
}
