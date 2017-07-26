package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copyJsonToSimpleValue)
	for _, kind := range []reflect.Kind{reflect.Int, reflect.String} {
		fromJsonMap[kind] = "CopyJsonToSimpleValue"
	}
}

var copyJsonToSimpleValue = generic.DefineFunc(
	"CopyJsonToSimpleValue(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators("opFuncName", genOpFuncName).
	Source(`
*dst = src.Read{{.DT|elem|opFuncName}}()
	`)
