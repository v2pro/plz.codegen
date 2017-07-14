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
	defDst := *dst
	dstLen := len(defDst)
	if len(src) < dstLen {
		dstLen = len(src)
	}
	for i := 0; i < dstLen; i++ {
		{{ $cp.FuncName }}(&defDst[i], src[i])
	}
	for i := dstLen; i < len(src); i++ {
		newElem := new({{ .DT|ptrSliceElem|elem|name }})
		{{ $cp.FuncName }}(newElem, src[i])
		defDst = append(defDst, *newElem)
	}
	*dst = defDst
	return nil
}`,
	FuncMap: map[string]interface{}{
		"ptrSliceElem": funcPtrSliceElem,
	},
}

func funcPtrSliceElem(typ reflect.Type) reflect.Type {
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
