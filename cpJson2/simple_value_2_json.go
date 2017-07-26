package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	cp2.Anything.ImportFunc(copySimpleValueToJson)
	for _, kind := range []reflect.Kind{reflect.Int} {
		toJsonMap[kind] = "CopySimpleValueToJson"
	}
}

var copySimpleValueToJson = generic.DefineFunc(
	"CopySimpleValueToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators("opFuncName", genOpFuncName).
	Source(`
dst.Write{{.ST|opFuncName}}(src)
	`)

func genOpFuncName(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	case reflect.Int16:
		return "Int16"
	case reflect.Int32:
		return "Int32"
	case reflect.Int64:
		return "Int64"
	case reflect.Uint:
		return "Uint"
	case reflect.Uint8:
		return "Uint8"
	case reflect.Uint16:
		return "Uint16"
	case reflect.Uint32:
		return "Uint32"
	case reflect.Uint64:
		return "Uint64"
	case reflect.Float32:
		return "Float32"
	case reflect.Float64:
		return "Float64"
	case reflect.String:
		return "String"
	case reflect.Bool:
		return "Bool"
	}
	return ""
}
