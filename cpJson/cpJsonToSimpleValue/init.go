package cpJsonToSimpleValue

import (
	"github.com/v2pro/wombat/cp/cpAnything"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cpJson/cpSimpleValueToJson"
)

func init() {
	cpAnything.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	TemplateName: "cpJsonToSimpleValue",
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	*dst = src.Read{{ .DT|elem|opFuncName }}()
}
`,
	GenMap: map[string]interface{}{
		"opFuncName": cpSimpleValueToJson.GenOpFuncName,
	},
}