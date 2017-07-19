package cpSliceToJson

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpAnything"
)

func init() {
	cpAnything.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	TemplateName: "cpSliceToJson",
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
	dst.WriteArrayStart()
	for i, elem := range src {
		if i != 0 {
			dst.WriteMore()
		}
		{{ $cpElem }}(err, dst, elem)
	}
	dst.WriteArrayEnd()
}
`,
}