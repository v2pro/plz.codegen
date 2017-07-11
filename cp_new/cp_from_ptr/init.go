package cp_from_ptr

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp_new/cp_statically"
)

func init() {
	cp_statically.F.Dependencies["cp_from_ptr"] = F
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cp_statically": cp_statically.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cp_statically" "DT" .DT "ST" (.ST|elem) }}
{{ $cp.Source }}
func {{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	if src == nil {
		return nil
	}
	return typed_{{ .funcName }}(
		dst.({{ .DT|name }}),
		src.({{ .ST|name }}))
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
