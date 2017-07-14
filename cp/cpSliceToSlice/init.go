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
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cpStatically" "DT" (.DT|ptrSliceElem) "ST" (.ST|elem) }}
{{ $cp.Source }}
func {{ .funcName }}(
	dst {{ .DT|name }},
	src {{ .ST|name }}) error {
	// end of signature
	dstLen := len(*dst)
	if len(src) < dstLen {
		dstLen = len(src)
	}
	for i := 0; i < dstLen; i++ {
		{{ $cp.FuncName }}(&(*dst)[i], src[i])
	}
	{{ if .DT|isSlice }}
	defDst := *dst
	for i := dstLen; i < len(src); i++ {
		newElem := new({{ .DT|ptrSliceElem|elem|name }})
		{{ $cp.FuncName }}(newElem, src[i])
		defDst = append(defDst, *newElem)
	}
	*dst = defDst
	{{ end }}
	return nil
}`,
	FuncMap: map[string]interface{}{
		"ptrSliceElem": funcPtrSliceElem,
		"isSlice": funcIsSlice,
	},
}

func funcPtrSliceElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	typ = typ.Elem()
	if typ.Kind() != reflect.Slice && typ.Kind() != reflect.Array {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}

func funcIsSlice(typ reflect.Type) bool {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	return typ.Kind() == reflect.Slice
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
