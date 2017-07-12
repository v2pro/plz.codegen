package cpFromPtr

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
)

func init() {
	cpStatically.F.Dependencies["cpFromPtr"] = F
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cpStatically": cpStatically.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cpStatically" "DT" .DT "ST" (.ST|elem) }}
{{ $cp.Source }}
func {{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	if dst == nil {
		return nil
	}
	if src == nil {
		return nil
	}
	return typed_{{ .funcName }}(
		{{ cast "dst" .DT }},
		{{ cast "src" .ST }})
}
func typed_{{ .funcName }}(
	dst {{ .DT|name }},
	src {{ .ST|name }}) error {
	// end of signature
	if src == nil {
		return nil
	}
	return typed_{{ $cp.FuncName }}(dst, *src)
}
`,
}
