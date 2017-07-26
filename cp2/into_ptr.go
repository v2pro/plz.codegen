package cp2

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copyIntoPtr)
}

var copyIntoPtr = generic.DefineFunc("CopyIntoPtr(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"isMap", func(typ reflect.Type) bool {
		return typ.Kind() == reflect.Map
	},
	"isNullable", func(typ reflect.Type) bool {
		switch typ.Kind() {
		case reflect.Ptr, reflect.Map, reflect.Interface, reflect.Slice:
			return true
		}
		return false
	}).
	Source(`
{{ $cp := expand "CopyAnything" "DT" (.DT|elem) "ST" .ST }}
{{ if .ST|isNullable }}
	if src == nil {
		*dst = nil
		return
	}
{{ end }}
defDst := *dst
if defDst == nil {
	{{ if .DT|elem|isMap }}
		defDst = {{ .DT|elem|name }}{}
	{{ else }}
		defDst = new({{ .DT|elem|elem|name }})
	{{ end }}
	{{$cp}}(err, defDst, src)
	*dst = defDst
	return
}
{{$cp}}(err, *dst, src)
`)
