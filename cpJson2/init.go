package cpJson2

import (
	"reflect"
	"github.com/json-iterator/go"
	"github.com/v2pro/wombat/cp2"
)

var iteratorType = reflect.TypeOf(new(jsoniter.Iterator))
var streamType = reflect.TypeOf(new(jsoniter.Stream))
var toJsonMap = map[reflect.Kind]string{}
var fromJsonMap = map[reflect.Kind]string{}

func init() {
	cp2.Dispatchers = append(cp2.Dispatchers, func(dstType, srcType reflect.Type) string {
		if srcType != iteratorType {
			return ""
		}
		if dstType.Kind() == reflect.Map {
			return "CopyJsonToMap"
		}
		if dstType.Kind() != reflect.Ptr {
			return ""
		}
		return fromJsonMap[dstType.Elem().Kind()]
	})
	cp2.Dispatchers = append(cp2.Dispatchers, func(dstType, srcType reflect.Type) string {
		if dstType != streamType {
			return ""
		}
		return toJsonMap[srcType.Kind()]
	})
}
