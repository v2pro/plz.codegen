package cp

import (
	// register functions into cpAnything
	_ "github.com/v2pro/wombat/cp/cpArrayToArray"
	_ "github.com/v2pro/wombat/cp/cpFromInterface"
	_ "github.com/v2pro/wombat/cp/cpFromPtr"
	_ "github.com/v2pro/wombat/cp/cpIntoInterface"
	_ "github.com/v2pro/wombat/cp/cpIntoPtr"
	_ "github.com/v2pro/wombat/cp/cpMapToMap"
	_ "github.com/v2pro/wombat/cp/cpMapToStruct"
	_ "github.com/v2pro/wombat/cp/cpSimpleValue"
	_ "github.com/v2pro/wombat/cp/cpSliceToSlice"
	"github.com/v2pro/wombat/cp/cpAnything"
	_ "github.com/v2pro/wombat/cp/cpAnything"
	_ "github.com/v2pro/wombat/cp/cpStructToMap"
	_ "github.com/v2pro/wombat/cp/cpStructToStruct"
	"reflect"
)

func init() {
	cpAnything.Dispatchers = append(cpAnything.Dispatchers, dispatch)
}

// Gen generates a instance of cpAnything
func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	return cpAnything.Gen(dstType, srcType)
}

func dispatch(dstType, srcType reflect.Type) string {
	if srcType.Kind() == reflect.Ptr {
		return "cpFromPtr"
	}
	if srcType.Kind() == reflect.Interface && srcType.NumMethod() == 0 {
		return "cpFromInterface"
	}
	if dstType.Kind() == reflect.Map &&
		srcType.Kind() == reflect.Map {
		return "cpMapToMap"
	}
	if dstType.Kind() == reflect.Map &&
		srcType.Kind() == reflect.Struct {
		return "cpStructToMap"
	}
	if dstType.Kind() == reflect.Ptr {
		if isSimpleValue(dstType.Elem()) && dstType.Elem().Kind() == srcType.Kind() {
			return "cpSimpleValue"
		}
		switch dstType.Elem().Kind() {
		case reflect.Interface:
			return "cpIntoInterface"
		case reflect.Ptr, reflect.Map:
			return "cpIntoPtr"
		case reflect.Struct:
			if srcType.Kind() == reflect.Struct {
				return "cpStructToStruct"
			}
			if srcType.Kind() == reflect.Map {
				return "cpMapToStruct"
			}
		case reflect.Slice:
			if srcType.Kind() == reflect.Slice {
				return "cpSliceToSlice"
			}
			if srcType.Kind() == reflect.Array {
				return "cpSliceToSlice"
			}
		case reflect.Array:
			if srcType.Kind() == reflect.Array {
				return "cpArrayToArray"
			}
			if srcType.Kind() == reflect.Slice {
				return "cpSliceToSlice"
			}
		}
	}
	return ""
}


func isSimpleValue(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:
		return true
	}
	return false
}