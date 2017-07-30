package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func init() {
	cp.Anything.ImportFunc(copyMapToJson)
	toJsonMap[reflect.Map] = "CopyMapToJson"
}

var copyMapToJson = generic.DefineFunc(
	"CopyMapToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" .DT "ST" (.ST|elem) }}
if src == nil {
	dst.WriteNil()
	return
}
dst.WriteObjectStart()
isFirst := true
for k, v := range src {
	if isFirst {
		isFirst = false
	} else {
		dst.WriteMore()
	}
	dst.WriteString(k)
	dst.WriteRaw(":")
	{{$cpElem}}(err, dst, v)
}
dst.WriteObjectEnd()
`)
