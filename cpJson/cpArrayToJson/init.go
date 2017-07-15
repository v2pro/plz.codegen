package cpArrayToJson

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
)

func init() {
	cpStatically.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cpArrayToJson",
	Dependencies: []*gen.FuncTemplate{cpStatically.F},
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
	dst.WriteArrayStart()
	{{ range $index, $_ := .ST|elems }}
		{{ if ne $index 0 }}
		dst.WriteMore()
		{{ end }}
		{{ $cpElem.FuncName }}(err, dst, src[{{ $index }}])
	{{ end }}
	dst.WriteArrayEnd()
}
`,
	GenMap: map[string]interface{}{
		"elems": genElems,
	},
}

func genElems(typ reflect.Type) []bool {
	return make([]bool, typ.Len())
}
