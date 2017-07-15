package cpSliceToJson

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpStatically"
)

func init() {
	cpStatically.F.Dependencies["cpSliceToJson"] = F
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
{{ $cpElem := gen "cpStatically" "DT" .DT "ST" (.ST|elem) }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	dst.WriteArrayStart()
	for i, elem := range src {
		if i != 0 {
			dst.WriteMore()
		}
		{{ $cpElem.FuncName }}(err, dst, elem)
	}
	dst.WriteArrayEnd()
}
`,
}