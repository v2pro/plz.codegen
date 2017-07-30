package cp

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copyArrayToArray)
}

var copyArrayToArray = generic.DefineFunc("CopyArrayToArray(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"minLength", func(dstType, srcType reflect.Type) int {
		if dstType.Elem().Len() < srcType.Len() {
			return dstType.Elem().Len()
		} else {
			return srcType.Len()
		}
	},
	"ptrArrayElem", func(dstType reflect.Type) reflect.Type {
		return reflect.PtrTo(dstType.Elem().Elem())
	}).
	Source(`
{{ $cp := expand "CopyAnything" "DT" (.DT|ptrArrayElem) "ST" (.ST|elem) }}
for i := 0; i < {{minLength .DT .ST}}; i++ {
	{{$cp}}(err, &dst[i], src[i])
}
`)
