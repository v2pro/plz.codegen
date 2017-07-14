package cpSimpleValue

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
)

func init() {
	cpStatically.F.Dependencies["cpSimpleValue"] = F
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
	//"cpSimpleValue": F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
func {{ .funcName }}(
	dst {{ .DT|name }},
	src {{ .ST|name }}) error {
	// end of signature
	*dst = src
	return nil
}
`,
}
