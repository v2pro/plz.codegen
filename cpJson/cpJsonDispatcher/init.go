package cpJsonDispatcher

import (
	"reflect"
	"github.com/json-iterator/go"
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
)

var jsoniterStreamType = reflect.TypeOf((*jsoniter.Stream)(nil))
var jsoniterIteratorType = reflect.TypeOf((*jsoniter.Iterator)(nil))

func init() {
	gen.ImportPackages["github.com/json-iterator/go"] = true
	cpStatically.Dispatchers = append(cpStatically.Dispatchers, dispatch)
}

func dispatch(dstType, srcType reflect.Type) string {
	if srcType == jsoniterIteratorType {
		if dstType.Kind() != reflect.Ptr {
			return ""
		}
		if GenOpFuncName(dstType.Elem()) != "" {
			return "cpJsonToSimpleValue"
		}
		if dstType.Elem().Kind() == reflect.Slice {
			return "cpJsonToSlice"
		}

	}
	if dstType == jsoniterStreamType {
		if GenOpFuncName(srcType) != "" {
			return "cpSimpleValueToJson"
		}
		if srcType.Kind() == reflect.Slice {
			return "cpSliceToJson"
		}
	}
	return ""
}

func GenOpFuncName(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	}
	return ""
}
