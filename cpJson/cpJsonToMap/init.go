package cpJsonToMap

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpJsonToMap"] = F
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cpStatically": cpStatically.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cpElem := gen "cpStatically" "DT" (.DT|ptrMapElem) "ST" .ST }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	src.ReadMapCB(func(iter *jsoniter.Iterator, key string) bool {
		elem, found := dst[key]
		if found {
			{{ $cpElem.FuncName }}(err, &elem, iter)
			dst[key] = elem
		} else {
			newElem := new({{ .DT|elem|name }})
			{{ $cpElem.FuncName }}(err, newElem, iter)
			dst[key] = *newElem
		}
		return true
	})
}
`,
	FuncMap: map[string]interface{}{
		"ptrMapElem": genPtrMapElem,
	},
}

func genPtrMapElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("expect map but found: " + typ.String())
	}
	return reflect.PtrTo(typ.Elem())
}