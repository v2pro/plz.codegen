package cpJson

import (
	_ "github.com/v2pro/wombat/cp"
	_ "github.com/v2pro/wombat/cpJson/cpSimpleValueToJson"
	_ "github.com/v2pro/wombat/cpJson/cpSliceToJson"
	_ "github.com/v2pro/wombat/cpJson/cpStructToJson"
	_ "github.com/v2pro/wombat/cpJson/cpMapToJson"
	_ "github.com/v2pro/wombat/cpJson/cpPtrToJson"
	_ "github.com/v2pro/wombat/cpJson/cpJsonToSimpleValue"
	_ "github.com/v2pro/wombat/cpJson/cpJsonToSlice"
	_ "github.com/v2pro/wombat/cpJson/cpJsonToStruct"
	_ "github.com/v2pro/wombat/cpJson/cpJsonToMap"
	_ "github.com/v2pro/wombat/cpJson/cpJsonToPtr"
	_ "github.com/v2pro/wombat/cpJson/cpJsonToArray"
	"reflect"
	"github.com/json-iterator/go"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/cpJson/cpSimpleValueToJson"
	"github.com/v2pro/plz"
)

var logger = plz.LoggerOf("package", "cpJson")
var jsoniterStreamType = reflect.TypeOf((*jsoniter.Stream)(nil))
var jsoniterIteratorType = reflect.TypeOf((*jsoniter.Iterator)(nil))

func init() {
	gen.ImportPackages["github.com/json-iterator/go"] = true
	cpStatically.Dispatchers = append(cpStatically.Dispatchers, dispatch)
}

func dispatch(dstType, srcType reflect.Type) string {
	if srcType == jsoniterIteratorType {
		if dstType.Kind() == reflect.Map {
			return "cpJsonToMap"
		}
		if dstType.Kind() != reflect.Ptr {
			return ""
		}
		if cpSimpleValueToJson.GenOpFuncName(dstType.Elem()) != "" {
			return "cpJsonToSimpleValue"
		}
		switch dstType.Elem().Kind() {
		case reflect.Slice:
			return "cpJsonToSlice"
		case reflect.Struct:
			return "cpJsonToStruct"
		case reflect.Map, reflect.Ptr:
			return "cpJsonToPtr"
		case reflect.Array:
			return "cpJsonToArray"
		}

	}
	if dstType == jsoniterStreamType {
		if cpSimpleValueToJson.GenOpFuncName(srcType) != "" {
			return "cpSimpleValueToJson"
		}
		switch srcType.Kind() {
		case reflect.Slice:
			return "cpSliceToJson"
		case reflect.Struct:
			return "cpStructToJson"
		case reflect.Map:
			return "cpMapToJson"
		case reflect.Array:
			return "cpSliceToJson"
		case reflect.Ptr:
			return "cpPtrToJson"
		}
	}
	return ""
}
