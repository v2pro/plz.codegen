package cpArrayToArray

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpArrayToArray"] = F
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
{{ $cpElem := gen "cpStatically" "DT" (.DT|ptrArrayElem) "ST" (.ST|elem) }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	for i := 0; i < {{ minLength .DT .ST }}; i++ {
		{{ $cpElem.FuncName }}(err, &dst[i], src[i])
	}
}`,
	GenMap: map[string]interface{}{
		"ptrArrayElem": genPtrArrayElem,
		"minLength":    genMinLength,
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

func genMinLength(dstType reflect.Type, srcType reflect.Type) int {
	len := dstType.Elem().Len()
	if srcType.Len() < len {
		len = srcType.Len()
	}
	return len
}
