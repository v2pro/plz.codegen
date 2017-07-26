package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copySliceToJson)
	toJsonMap[reflect.Slice] = "CopySliceToJson"
}

var copySliceToJson = generic.DefineFunc(
	"CopySliceToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" .DT "ST" (.ST|elem) }}
if src == nil {
	dst.WriteNil()
	return
}
dst.WriteArrayStart()
for i, elem := range src {
	if i != 0 {
		dst.WriteMore()
	}
	{{$cpElem}}(err, dst, elem)
}
dst.WriteArrayEnd()
`)
