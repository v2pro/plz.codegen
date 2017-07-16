package maxSimpleValue

import (
	"github.com/v2pro/plz/util"
	"github.com/v2pro/wombat/fp/cmpSimpleValue"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	util.GenMaxSimpleValue = genF
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "maxSimpleValue",
	Dependencies: []*gen.FuncTemplate{cmpSimpleValue.F},
	TemplateParams: map[string]string{
		"T": "the type to max",
	},
	FuncName: `Max_{{ .T|symbol }}`,
	Source: `
{{ $compare := gen "cmpSimpleValue" "T" .T }}
func Exported_{{ .funcName }}(objs []interface{}) interface{} {
	currentMax := objs[0].({{ .T|name }})
	for i := 1; i < len(objs); i++ {
		typedObj := objs[i].({{ .T|name }})
		if {{ $compare }}(typedObj, currentMax) > 0 {
			currentMax = typedObj
		}
	}
	return currentMax
}
func {{ .funcName }}(objs []{{ .T|name }}) {{ .T|name }} {
	currentMax := objs[0]
	for i := 1; i < len(objs); i++ {
		if {{ $compare }}(objs[i], currentMax) > 0 {
			currentMax = objs[i]
		}
	}
	return currentMax
}`,
}

func genF(typ reflect.Type) func([]interface{}) interface{} {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8:
		funcObj := gen.Expand(F, `T`, typ)
		return funcObj.(func([]interface{}) interface{})
	}
	return nil
}
