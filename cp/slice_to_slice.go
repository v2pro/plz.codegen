package cp

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copySliceToSlice)
}

var copySliceToSlice = generic.DefineFunc("CopySliceToSlice(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"ptrSliceElem", func(dstType reflect.Type) reflect.Type {
		return reflect.PtrTo(dstType.Elem().Elem())
	}).
	Source(`
{{ $cp := expand "CopyAnything" "DT" (.DT|ptrSliceElem) "ST" (.ST|elem) }}
if src == nil {
	*dst = nil
	return
}
dstLen := len(*dst)
if len(src) < dstLen {
	dstLen = len(src)
}
for i := 0; i < dstLen; i++ {
	{{$cp}}(err, &(*dst)[i], src[i])
}
defDst := *dst
for i := dstLen; i < len(src); i++ {
	newElem := new({{ .DT|elem|elem|name }})
	{{$cp}}(err, newElem, src[i])
	defDst = append(defDst, *newElem)
}
*dst = defDst
`)
