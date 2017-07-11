package max

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/fp/compare"
	"reflect"
)

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"compareFuncName": compare.F,
	},
	Variables: map[string]string{
		"T": "the type to max",
	},
	FuncName: `Max_{{ .T|name }}`,
	Source: `
func {{ .funcName }}(objs []interface{}) interface{} {
	currentMax := objs[0].({{ .T|name }})
	for i := 1; i < len(objs); i++ {
		typedObj := objs[i].({{ .T|name }})
		if {{ .compareFuncName }}(typedObj, currentMax) > 0 {
			currentMax = typedObj
		}
	}
	return currentMax
}
func typed_{{ .funcName }}(objs []{{ .T|name }}) {{ .T|name }} {
	currentMax := objs[0]
	for i := 1; i < len(objs); i++ {
		if {{ .compareFuncName }}(objs[i], currentMax) > 0 {
			currentMax = objs[i]
		}
	}
	return currentMax
}`,
}

var symbols = map[reflect.Type]func([]interface{}) interface{}{}

func Call(objs []interface{}) interface{} {
	typ := reflect.TypeOf(objs[0])
	f := symbols[typ]
	if f == nil {
		funcObj := gen.Compile(F, `T`, typ)
		f = funcObj.(func([]interface{}) interface{})
		symbols[typ] = f
	}
	return f(objs)
}