package cp

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copyMapToMap)
}

var copyMapToMap = generic.DefineFunc("CopyMapToMap(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"ptrMapElem", func(typ reflect.Type) reflect.Type {
		if typ.Kind() != reflect.Map {
			panic("expects Map, but found " + typ.String())
		}
		return reflect.PtrTo(typ.Elem())
	},
	"ptrMapKey", func(typ reflect.Type) reflect.Type {
		if typ.Kind() != reflect.Map {
			panic("expects Map, but found " + typ.String())
		}
		return reflect.PtrTo(typ.Key())
	},
	"mapKey", func(typ reflect.Type) reflect.Type {
		if typ.Kind() != reflect.Map {
			panic("expects Map, but found " + typ.String())
		}
		return typ.Key()
	}).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" (.DT|ptrMapElem) "ST" (.ST|elem) }}
{{ $cpKey := expand "CopyAnything" "DT" (.DT|ptrMapKey) "ST" (.ST|mapKey) }}
for key, elem := range src {
	existingElem, found := dst[key]
	if found {
		{{$cpElem}}(err, &existingElem, elem)
		dst[key] = existingElem
	} else {
		newKey := new({{.DT|mapKey|name}})
		{{$cpKey}}(err, newKey, key)
		newElem := new({{.DT|elem|name}})
		{{$cpElem}}(err, newElem, elem)
		dst[*newKey] = *newElem
	}
}`)
