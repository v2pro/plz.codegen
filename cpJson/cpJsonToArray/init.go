package cpJsonToArray

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
	TemplateName: "cpJsonToArray",
	Dependencies: []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cpElem := gen "cpAnything" "DT" (.DT|ptrArrayElem) "ST" .ST }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	index := 0
	src.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		if index < {{ .DT|arrayLen }} {
			{{ $cpElem }}(err, &((*dst)[index]), iter)
		} else {
			iter.Skip()
		}
		index++
		return true
	})
}
`,
	GenMap: map[string]interface{}{
		"ptrArrayElem": genPtrArrayElem,
		"arrayLen":     genArrayLen,
	},
}

func genPtrArrayElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	typ = typ.Elem()
	if typ.Kind() != reflect.Array {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}

func genArrayLen(typ reflect.Type) int {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	typ = typ.Elem()
	if typ.Kind() != reflect.Array {
		panic("unexpected")
	}
	return typ.Len()
}
