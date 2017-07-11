package cp_statically

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
	},
	FuncMap: map[string]interface{}{
		"mustBeWritable": func_mustBeWritable,
		"isDirectPtr": func_isDirectPtr,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ .DT|mustBeWritable }}
{{ if .ST|isPtr }}
	{{ $cp := gen "cp_from_ptr" "DT" .DT "ST" .ST }}
	{{ $cp.Source }}
{{ else }}
	{{ if .DT|isPtr }}
		{{ if .DT|isDirectPtr }}
			{{ $cp := gen "cp_simple_value" "DT" .DT "ST" .ST }}
			{{ $cp.Source }}
		{{ else }}
			{{ $cp := gen "cp_into_ptr" "DT" .DT "ST" .ST }}
			{{ $cp.Source }}
		{{ end }}
	{{ end }} {{/* .DT|isPtr */}}
{{ end }} {{/* .ST|isPtr */}}
`,
}

func func_mustBeWritable(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Map:
		return ""
	}
	panic("destination type is not writable: " + typ.String())
}

func func_isDirectPtr(typ reflect.Type) bool {
	if typ.Kind() != reflect.Ptr {
		return false
	}
	return typ.Elem().Kind() != reflect.Ptr
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
