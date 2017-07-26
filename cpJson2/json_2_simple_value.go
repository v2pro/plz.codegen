package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copyJsonToSimpleValue)
	for kind := range simpleValueMap {
		fromJsonMap[kind] = "CopyJsonToSimpleValue"
	}
}

var copyJsonToSimpleValue = generic.DefineFunc(
	"CopyJsonToSimpleValue(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators(
	"opFuncName", func(typ reflect.Type) string {
		return simpleValueMap[typ.Kind()]
	}).
	Source(`
*dst = src.Read{{.DT|elem|opFuncName}}()
	`)
