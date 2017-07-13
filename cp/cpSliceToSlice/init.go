package cpSliceToSlice

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpSliceToSlice"] = F
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
{{ $cp := gen "cpStatically" "DT" (.DT|sliceElem) "ST" (.ST|elem) }}
{{ $cp.Source }}
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
	defDst := *dst
	for i := 0; i < len(src); i++ {
		defDst = append(defDst, src[i])
	}
	*dst = defDst
	return nil
}`,
	FuncMap: map[string]interface{}{
		"sliceElem": funcSliceElem,
	},
}

func funcSliceElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	typ = typ.Elem()
	if typ.Kind() != reflect.Slice {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
