package max

import (
	"github.com/v2pro/plz/util"
	"github.com/v2pro/wombat/fp/cmpStructByField"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	util.GenMaxStructByField = genF
}

// F the function definition
var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
		"cmpStructByField": cmpStructByField.F,
	},
	TemplateParams: map[string]string{
		"T": "the struct type to max",
		"F": "the field name of T",
	},
	FuncName: `Max_{{ .T|name }}_by_{{ .F }}`,
	Source: `
{{ $compare := gen "cmpStructByField" "T" .T "F" .F }}
{{ $compare.Source }}
func Exported_{{ .funcName }}(objs []interface{}) interface{} {
	currentMaxObj := objs[0]
	for i := 1; i < len(objs); i++ {
		currentMax := {{ cast "currentMaxObj" .T }}
		elem := {{ cast "objs[i]" .T }}
		if {{ $compare.FuncName }}(elem, currentMax) > 0 {
			currentMaxObj = objs[i]
		}
	}
	return currentMaxObj
}`,
}

func genF(typ reflect.Type, fieldName string) func([]interface{}) interface{} {
	funcObj := gen.Compile(F, "T", typ, "F", fieldName)
	return funcObj.(func([]interface{}) interface{})
}
