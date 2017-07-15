package cpJsonToPtr

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpJsonToPtr"] = F
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
{{ $cp := gen "cpStatically" "DT" (.DT|elem) "ST" .ST }}
{{ $cp.Source }}
// generated from cpIntoPtr
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	if src.ReadNil() {
		return
	}
	if dst == nil {
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
	FuncMap: map[string]interface{}{
		"isMap": funcIsMap,
	},
}

func funcIsMap(typ reflect.Type) bool {
	return typ.Kind() == reflect.Map
}