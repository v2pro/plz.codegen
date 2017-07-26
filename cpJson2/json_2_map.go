package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copyJsonToMap)
	// dispatch handled in init.go directly
}

var copyJsonToMap = generic.DefineFunc(
	"CopyJsonToMap(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators(
	"ptrMapElem", func(typ reflect.Type) reflect.Type {
		return reflect.PtrTo(typ.Elem())
	}).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" (.DT|ptrMapElem) "ST" .ST }}
src.ReadMapCB(func(iter *jsoniter.Iterator, key string) bool {
	elem, found := dst[key]
	if found {
		{{$cpElem}}(err, &elem, iter)
		dst[key] = elem
	} else {
		newElem := new({{.DT|elem|name}})
		{{$cpElem}}(err, newElem, iter)
		dst[key] = *newElem
	}
	return true
})`)
