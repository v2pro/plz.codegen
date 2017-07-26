package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copyJsonToArray)
	fromJsonMap[reflect.Array] = "CopyJsonToArray"
}

var copyJsonToArray = generic.DefineFunc(
	"CopyJsonToArray(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators(
	"ptrArrayElem", func(typ reflect.Type) reflect.Type {
		return reflect.PtrTo(typ.Elem().Elem())
	},
	"arrayLen", func(typ reflect.Type) int {
		return typ.Elem().Len()
	}).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" (.DT|ptrArrayElem) "ST" .ST }}
index := 0
src.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
	if index < {{.DT|arrayLen}} {
		{{$cpElem}}(err, &dst[index], iter)
	} else {
		iter.Skip()
	}
	index++
	return true
})
	`)
