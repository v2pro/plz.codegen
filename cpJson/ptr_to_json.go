package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func init() {
	cp.Anything.ImportFunc(copyPtrToJson)
	toJsonMap[reflect.Ptr] = "CopyPtrToJson"
}

var copyPtrToJson = generic.DefineFunc("CopyPtrToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
	Source(`
{{ $cp := expand "CopyAnything" "DT" .DT "ST" (.ST|elem) }}
if src == nil {
	dst.WriteNil()
	return
}
{{$cp}}(err, dst, *src)
`)
