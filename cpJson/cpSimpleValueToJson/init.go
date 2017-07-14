package cpSimpleValueToJson

import (
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"reflect"
	"github.com/json-iterator/go"
)

type streamInterface interface {
	WriteInt(val int)
	WriteInt8(val int8)
}

var jsoniterStreamType = reflect.TypeOf((*jsoniter.Stream)(nil))
var streamInterfaceType = reflect.TypeOf((*streamInterface)(nil)).Elem()

func init() {
	cpStatically.F.Dependencies["cpSimpleValueToJson"] = F
	cpStatically.Dispatchers = append(cpStatically.Dispatchers, dispatchJson)
	gen.TypeTranslator = append(gen.TypeTranslator, translateType)
}

func translateType(typ reflect.Type) reflect.Type {
	if typ == jsoniterStreamType {
		return streamInterfaceType
	}
	return typ
}

func dispatchJson(dstType, srcType reflect.Type) string {
	if dstType == streamInterfaceType {
		return "cpSimpleValueToJson"
	}
	return ""
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		//"cpSimpleValue": F,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	dst.Write{{ .ST|opFuncName }}(src)
}
`,
	FuncMap: map[string]interface{}{
		"opFuncName": genOpFuncName,
	},
}

func genOpFuncName(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	}
	panic("not implemented")
}
