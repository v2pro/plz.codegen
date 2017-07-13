package cpSliceToSlice

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpStatically.F.Dependencies["cpMapToMap"] = F
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
{{ $cpElem := gen "cpStatically" "DT" (.DT|ptrMapElem) "ST" (.ST|elem) }}
{{ $cpElem.Source }}
{{ $cpKey := gen "cpStatically" "DT" (.DT|ptrMapKey) "ST" (.ST|mapKey) }}
{{ $cpKey.Source }}
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
	for key, elem := range src {
		existingElem, found := dst[key]
		if found {
			typed_{{ $cpElem.FuncName }}(&existingElem, elem)
			dst[key] = existingElem
		} else {
			newKey := new({{ .DT|mapKey|name }})
			typed_{{ $cpKey.FuncName }}(newKey, key)
			newElem := new({{ .DT|elem|name }})
			typed_{{ $cpElem.FuncName }}(newElem, elem)
			dst[*newKey] = *newElem
		}
	}
	return nil
}`,
	FuncMap: map[string]interface{}{
		"ptrMapElem": funcPtrMapElem,
		"ptrMapKey": funcPtrMapKey,
		"mapKey": funcMapKey,
	},
}

func funcPtrMapElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}

func funcPtrMapKey(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Key())
}

func funcMapKey(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("unexpected")
	}
	return typ.Key()
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
