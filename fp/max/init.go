package max

import (
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/wombat/fp/compare"
	"reflect"
	"github.com/v2pro/plz/util"
)

func init() {
	util.GenMaxSimpleValue = Gen
}

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

func Gen(typ reflect.Type) func([]interface{}) interface{} {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8:
		funcObj := gen.Compile(F, `T`, typ)
		return funcObj.(func([]interface{}) interface{})
	}
	return nil
}
