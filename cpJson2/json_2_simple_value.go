package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
	"github.com/json-iterator/go"
)

var iteratorType = reflect.TypeOf(new(jsoniter.Iterator))

func init() {
	cp2.Anything.ImportFunc(copyJsonToSimpleValue)
	cp2.Dispatchers = append(cp2.Dispatchers, func(dstType, srcType reflect.Type) string {
		if srcType != iteratorType {
			return ""
		}
		if dstType.Kind() != reflect.Ptr {
			return ""
		}
		switch dstType.Elem().Kind() {
		case reflect.Int:
			return "CopyJsonToSimpleValue"
		}
		return ""
	})
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