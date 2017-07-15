package cpMapToMap

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
	FuncTemplateName: "cpMapToMap",
	Dependencies: []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cpElem := gen "cpAnything" "DT" (.DT|ptrMapElem) "ST" (.ST|elem) }}
{{ $cpKey := gen "cpAnything" "DT" (.DT|ptrMapKey) "ST" (.ST|mapKey) }}
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	for key, elem := range src {
		existingElem, found := dst[key]
		if found {
			{{ $cpElem }}(err, &existingElem, elem)
			dst[key] = existingElem
		} else {
			newKey := new({{ .DT|mapKey|name }})
			{{ $cpKey }}(err, newKey, key)
			newElem := new({{ .DT|elem|name }})
			{{ $cpElem }}(err, newElem, elem)
			dst[*newKey] = *newElem
		}
	}
	return
}`,
	GenMap: map[string]interface{}{
		"ptrMapElem": genPtrMapElem,
		"ptrMapKey":  genPtrMapKey,
		"mapKey":     genMapKey,
	},
}

func genPtrMapElem(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Elem())
}

func genPtrMapKey(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("unexpected")
	}
	return reflect.PtrTo(typ.Key())
}

func genMapKey(typ reflect.Type) reflect.Type {
	if typ.Kind() != reflect.Map {
		panic("unexpected")
	}
	return typ.Key()
}
