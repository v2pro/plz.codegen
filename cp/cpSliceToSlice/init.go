package cpSliceToSlice

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cpSliceToSlice",
	Dependencies: []*gen.FuncTemplate{cpStatically.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cpStatically" "DT" (.DT|ptrSliceElem) "ST" (.ST|elem) }}
{{ $cp.Source }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	dstLen := len(*dst)
	if len(src) < dstLen {
		dstLen = len(src)
	}
	for i := 0; i < dstLen; i++ {
		{{ $cp.FuncName }}(err, &(*dst)[i], src[i])
	}
	{{ if .DT|isSlice }}
	defDst := *dst
	for i := dstLen; i < len(src); i++ {
		newElem := new({{ .DT|ptrSliceElem|elem|name }})
		{{ $cp.FuncName }}(err, newElem, src[i])
		defDst = append(defDst, *newElem)
	}
	*dst = defDst
	{{ end }}
}`,
	GenMap: map[string]interface{}{
		"ptrSliceElem": genPtrSliceElem,
		"isSlice":      genIsSlice,
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

func genIsSlice(typ reflect.Type) bool {
	if typ.Kind() != reflect.Ptr {
		panic("unexpected")
	}
	return typ.Elem().Kind() == reflect.Slice
}
