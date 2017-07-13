package cpArrayToArray

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpArrayToArray"] = F
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cpStatically": cpStatically.F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cpElem := gen "cpStatically" "DT" (.DT|ptrArrayElem) "ST" (.ST|elem) }}
{{ $cpElem.Source }}
func {{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	return typed_{{ .funcName }}(
		{{ cast "dst" .DT }},
		{{ cast "src" .ST }})
}
func typed_{{ .funcName }}(
	dst {{ .DT|name }},
	src {{ .ST|name }}) error {
	// end of signature
	for i := 0; i < {{ minLength .DT .ST }}; i++ {
		err := typed_{{ $cpElem.FuncName }}(&dst[i], src[i])
		if err != nil {
			return err
		}
	}
	return nil
}`,
	FuncMap: map[string]interface{}{
		"ptrArrayElem": funcPtrArrayElem,
		"minLength":    funcMinLength,
	},
}

func funcPtrArrayElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	typ = typ.Elem()
	if typ.Kind() != reflect.Array {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}

func funcMinLength(dstType reflect.Type, srcType reflect.Type) int {
	len := dstType.Elem().Len()
	if srcType.Len() < len {
		len = srcType.Len()
	}
	return len
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
