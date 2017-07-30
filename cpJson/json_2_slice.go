package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func init() {
	cp.Anything.ImportFunc(copyJsonToSlice)
	fromJsonMap[reflect.Slice] = "CopyJsonToSlice"
}

var copyJsonToSlice = generic.DefineFunc(
	"CopyJsonToSlice(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
	Generators(
	"ptrSliceElem", func(typ reflect.Type) reflect.Type {
		return reflect.PtrTo(typ.Elem().Elem())
	}).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" (.DT|ptrSliceElem) "ST" .ST }}
if src.ReadNil() {
	*dst = nil
	return
}
index := 0
originalLen := len(*dst)
src.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
	if index < originalLen {
		elem := &(*dst)[index]
		{{$cpElem}}(err, elem, iter)
	} else {
		elem := new({{.DT|elem|elem|name}})
		{{$cpElem}}(err, elem, iter)
		*dst = append(*dst, *elem)
	}
	index++
	return true
})`)
