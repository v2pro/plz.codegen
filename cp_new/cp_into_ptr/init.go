package cp_into_ptr

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp_new/cp_statically"
)

func init() {
	cp_statically.F.Dependencies["cp_into_ptr"] = F
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
{{ $cp := gen "cp_statically" "DT" (.DT|elem) "ST" .ST }}
{{ $cp.Source }}
func {{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	if dst == nil {
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
	defDst := *dst
	if defDst == nil {
		defDst = new({{ .DT|elem|elem|name }})
		err := typed_{{ $cp.FuncName }}(defDst, src)
		*dst = defDst
		return err
	}
	return typed_{{ $cp.FuncName }}(*dst, src)
}
`,
}
