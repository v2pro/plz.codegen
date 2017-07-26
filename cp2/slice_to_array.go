package cp2

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copySliceToArray)
}

var copySliceToArray = generic.DefineFunc("CopySliceToArray(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"ptrArrayElem", func(dstType reflect.Type) reflect.Type {
		return reflect.PtrTo(dstType.Elem().Elem())
	}).
	Source(`
{{ $cp := expand "CopyAnything" "DT" (.DT|ptrArrayElem) "ST" (.ST|elem) }}
dstLen := len(*dst)
if len(src) < dstLen {
	dstLen = len(src)
}
for i := 0; i < dstLen; i++ {
	{{$cp}}(err, &dst[i], src[i])
}`)
