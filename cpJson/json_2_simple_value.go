package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func init() {
	cp.Anything.ImportFunc(copyJsonToSimpleValue)
	for kind := range simpleValueMap {
		fromJsonMap[kind] = "CopyJsonToSimpleValue"
	}
}

var copyJsonToSimpleValue = generic.DefineFunc(
	"CopyJsonToSimpleValue(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
	Generators(
	"opFuncName", func(typ reflect.Type) string {
		return simpleValueMap[typ.Kind()]
	}).
	Source(`
*dst = src.Read{{.DT|elem|opFuncName}}()
	`)
