package cpJsonToArray

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpJsonToArray"] = F
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
{{ $cpElem := gen "cpStatically" "DT" (.DT|ptrArrayElem) "ST" .ST }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	index := 0
	src.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		if index < {{ .DT|arrayLen }} {
			{{ $cpElem.FuncName }}(err, &((*dst)[index]), iter)
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
		"arrayLen": genArrayLen,
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