package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copyJsonToPtr)
	fromJsonMap[reflect.Map] = "CopyJsonToPtr"
	fromJsonMap[reflect.Ptr] = "CopyJsonToPtr"
}

var copyJsonToPtr = generic.DefineFunc(
	"CopyJsonToPtr(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators(
	"isMap", func(typ reflect.Type) bool {
		return typ.Kind() == reflect.Map
	}).
	Source(`
{{ $cp := expand "CopyAnything" "DT" (.DT|elem) "ST" .ST }}
if dst == nil {
	src.Skip()
	return
}
if src.ReadNil() {
	*dst = nil
	return
}
defDst := *dst
if defDst == nil {
	{{ if .DT|elem|isMap }}
		defDst = {{.DT|elem|name}}{}
	{{ else }}
		defDst = new({{.DT|elem|elem|name}})
	{{ end }}
	{{$cp}}(err, defDst, src)
	*dst = defDst
	return
}
{{ $cp }}(err, *dst, src)
`)
