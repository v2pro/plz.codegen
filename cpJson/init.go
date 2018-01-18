package cpJson

import (
	"reflect"
	"github.com/json-iterator/go"
	"github.com/v2pro/wombat/cp"
)

var iteratorType = reflect.TypeOf(new(jsoniter.Iterator))
var streamType = reflect.TypeOf(new(jsoniter.Stream))
var toJsonMap = map[reflect.Kind]string{}
var fromJsonMap = map[reflect.Kind]string{}
var simpleValueMap = map[reflect.Kind]string{
	reflect.Int: "Int",
	reflect.Int8: "Int8",
	reflect.Int16: "Int16",
	reflect.Int32: "Int32",
	reflect.Int64: "Int64",
	reflect.Uint: "Uint",
	reflect.Uint8: "Uint8",
	reflect.Uint16: "Uint16",
	reflect.Uint32: "Uint32",
	reflect.Uint64: "Uint64",
	reflect.Float32: "Float32",
	reflect.Float64: "Float64",
	reflect.String: "StringValue",
	reflect.Bool: "BoolValue",
}

func init() {
	cp.Dispatchers = append(cp.Dispatchers, func(dstType, srcType reflect.Type) string {
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
	cp.Dispatchers = append(cp.Dispatchers, func(dstType, srcType reflect.Type) string {
		if dstType != streamType {
			return ""
		}
		return toJsonMap[srcType.Kind()]
	})
}
