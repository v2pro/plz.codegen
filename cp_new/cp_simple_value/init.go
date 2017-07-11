package cp_simple_value

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp_new/cp_statically"
)

func init() {
	cp_statically.F.Dependencies["cp_simple_value"] = F
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		//"cp_simple_value": F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
func {{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	if dst == nil {
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
	*dst = src
	return nil
}
`,
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
