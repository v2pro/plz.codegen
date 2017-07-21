package max

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/compare"
	"reflect"
	"github.com/v2pro/plz/util"
)

func init() {
	util.GenMaxByField = func(typ reflect.Type, fieldName string) func(collection []interface{}) interface{} {
		return generic.Expand(ByFieldForPlz, "T", typ, "F", fieldName).
		(func(collection []interface{}) interface{})
	}
}

var ByField = generic.DefineFunc("MaxByField(vals T) E").
	Param("T", "array type").
	Param("E", "array element type", func(argMap generic.ArgMap) interface{} {
	return argMap["T"].(reflect.Type).Elem()
}).
	Param("F", "the field to compare").
	ImportFunc(compare.ByField).
	Source(`
{{ $compare := expand "CompareByField" "T" .E "F" .F }}
currentMax := vals[0]
for i := 1; i < len(vals); i++ {
	if {{$compare}}(vals[i], currentMax) > 0 {
		currentMax = vals[i]
	}
}
return currentMax`)

var ByFieldForPlz = generic.DefineFunc("MaxByFieldForPlz(vals []interface{}) interface{}").
	Param("T", "array element type").
	Param("F", "the field to compare").
	ImportFunc(compare.ByField).
	Source(`
{{ $compare := expand "CompareByField" "T" .T "F" .F }}
currentMax := vals[0].({{.T|name}})
for i := 1; i < len(vals); i++ {
	typedVal := vals[i].({{.T|name}})
	if {{$compare}}(typedVal, currentMax) > 0 {
		currentMax = typedVal
	}
}
return currentMax`)
