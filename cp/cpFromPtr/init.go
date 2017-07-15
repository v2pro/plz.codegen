package cpFromPtr

import (
	"github.com/v2pro/wombat/cp/cpAnything"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpAnything.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cpFromPtr",
	Dependencies:     []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cpAnything" "DT" .DT "ST" (.ST|elem) }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	if dst == nil {
		return
	}
	if src == nil {
		{{ if .DT|isPtrNullable }}
		*dst = nil
		{{ end }}
		return
	}
	{{ $cp }}(err, dst, *src)
}
`,
	GenMap: map[string]interface{}{
		"isPtrNullable": genIsPtrNullable,
	},
}

func genIsPtrNullable(typ reflect.Type) bool {
	if typ.Kind() != reflect.Ptr {
		return false
	}
	switch typ.Elem().Kind() {
	case reflect.Ptr, reflect.Map, reflect.Interface, reflect.Slice:
		return true
	}
	return false
}
