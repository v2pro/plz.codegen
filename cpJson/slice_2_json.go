package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func init() {
	cp.Anything.ImportFunc(copySliceToJson)
	toJsonMap[reflect.Slice] = "CopySliceToJson"
}

var copySliceToJson = generic.DefineFunc(
	"CopySliceToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
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
