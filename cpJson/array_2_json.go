package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func init() {
	cp.Anything.ImportFunc(copyArrayToJson)
	toJsonMap[reflect.Array] = "CopyArrayToJson"
}

var copyArrayToJson = generic.DefineFunc(
	"CopyArrayToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
	Generators(
	"elems", func(typ reflect.Type) []bool {
		return make([]bool, typ.Len())
	}).
	Source(`
{{ $cpElem := expand "CopyAnything" "DT" .DT "ST" (.ST|elem) }}
dst.WriteArrayStart()
{{ range $index, $_ := .ST|elems }}
	{{ if ne $index 0 }}
	dst.WriteMore()
	{{ end }}
	{{$cpElem}}(err, dst, src[{{$index}}])
{{ end }}
dst.WriteArrayEnd()
`)
