package cpJsonToSimpleValue

import (
	"reflect"
	"github.com/json-iterator/go"
	"github.com/v2pro/wombat/cp/cpStatically"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/cpJson/cpSimpleValueToJson"
)

type jsoniterIterator interface {
	ReadInt() int
	ReadInt8() int8
}

var jsoniterIteratorType = reflect.TypeOf((*jsoniter.Iterator)(nil))
var iteratorInterfaceType = reflect.TypeOf((*jsoniterIterator)(nil)).Elem()

func init() {
	cpStatically.F.Dependencies["cpJsonToSimpleValue"] = F
	cpStatically.Dispatchers = append(cpStatically.Dispatchers, dispatch)
	gen.TypeTranslator = append(gen.TypeTranslator, translateType)
}

func translateType(typ reflect.Type) reflect.Type {
	if typ == jsoniterIteratorType {
		return iteratorInterfaceType
	}
	return typ
}

func dispatch(dstType, srcType reflect.Type) string {
	if srcType == iteratorInterfaceType {
		return "cpJsonToSimpleValue"
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
	*dst = src.Read{{ .DT|elem|opFuncName }}()
}
`,
	FuncMap: map[string]interface{}{
		"opFuncName": cpSimpleValueToJson.GenOpFuncName,
	},
}