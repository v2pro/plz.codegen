package cpMapToJson

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpStatically"
)

func init() {
	cpStatically.F.Dependencies["cpMapToJson"] = F
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
{{ $cpElem := gen "cpStatically" "DT" .DT "ST" (.ST|elem) }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	if src == nil {
		dst.WriteNil()
		return
	}
	dst.WriteObjectStart()
	isFirst := true
	for k, v := range src {
		if isFirst {
			isFirst = false
		} else {
			dst.WriteMore()
		}
		dst.WriteString(k)
		dst.WriteRaw(":")
		{{ $cpElem.FuncName }}(err, dst, v)
	}
	dst.WriteObjectEnd()
}
`}