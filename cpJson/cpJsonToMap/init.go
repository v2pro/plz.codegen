package cpJsonToMap

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
	TemplateName: "cpJsonToMap",
	Dependencies: []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cpElem := gen "cpAnything" "DT" (.DT|ptrMapElem) "ST" .ST }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	src.ReadMapCB(func(iter *jsoniter.Iterator, key string) bool {
		elem, found := dst[key]
		if found {
			{{ $cpElem }}(err, &elem, iter)
			dst[key] = elem
		} else {
			newElem := new({{ .DT|elem|name }})
			{{ $cpElem }}(err, newElem, iter)
			dst[key] = *newElem
		}
		return true
	})
}
`,
	GenMap: map[string]interface{}{
		"ptrMapElem": genPtrMapElem,
	},
}

func genPtrMapElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("expect map but found: " + typ.String())
	}
	return reflect.PtrTo(typ.Elem())
}
