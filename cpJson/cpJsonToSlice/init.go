package cpJsonToSlice

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cpJson/cpJsonDispatcher"
)

func init() {
	cpStatically.F.Dependencies["cpJsonToSlice"] = F
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
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	src.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		*dst = append(*dst, iter.Read{{ .DT|elem|elem|opFuncName }}())
		return true
	})
}
`,
	FuncMap: map[string]interface{}{
		"opFuncName": cpJsonDispatcher.GenOpFuncName,
	},
}