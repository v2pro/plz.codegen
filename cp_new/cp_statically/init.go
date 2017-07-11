package cp_statically

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ if .ST|isPtr }}
	{{ $cp := gen "cp_from_ptr" "DT" .DT "ST" .ST }}
	{{ $cp.Source }}
{{ else }}
	{{ $cp := gen "cp_simple_value" "DT" .DT "ST" .ST }}
	{{ $cp.Source }}
{{ end }}
`,
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
