package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copyMapToJson)
	toJsonMap[reflect.Map] = "CopyMapToJson"
}

var copyMapToJson = generic.DefineFunc(
	"CopyMapToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
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
