package cpMapToJson

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpAnything"
)

func init() {
	cpAnything.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	TemplateName: "cpMapToJson",
	Dependencies: []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cpElem := gen "cpAnything" "DT" .DT "ST" (.ST|elem) }}
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
		{{ $cpElem }}(err, dst, v)
	}
	dst.WriteObjectEnd()
}
`}