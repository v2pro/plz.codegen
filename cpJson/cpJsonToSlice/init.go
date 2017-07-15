package cpJsonToSlice

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpJsonToSlice"] = F
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
{{ $cpElem := gen "cpStatically" "DT" (.DT|ptrSliceElem) "ST" .ST }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	if src.ReadNil() {
		*dst = nil
		return
	}
	index := 0
	originalLen := len(*dst)
	src.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		if index < originalLen {
			elem := &(*dst)[index]
			{{ $cpElem.FuncName }}(err, elem, iter)
		} else {
			elem := new({{ .DT|elem|elem|name }})
			{{ $cpElem.FuncName }}(err, elem, iter)
			*dst = append(*dst, *elem)
		}
		index++
		return true
	})
}
`,
	GenMap: map[string]interface{}{
		"ptrSliceElem": genPtrSliceElem,
	},
}

func genPtrSliceElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	typ = typ.Elem()
	if typ.Kind() != reflect.Slice && typ.Kind() != reflect.Array {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}