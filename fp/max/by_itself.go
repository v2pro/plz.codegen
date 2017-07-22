package max

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/compare"
	"reflect"
	"github.com/v2pro/plz/util"
)

func init() {
	util.GenMaxByItself = func(typ reflect.Type) func(collection []interface{}) interface{} {
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return generic.Expand(ByItselfForPlz, "T", typ).
			(func(collection []interface{}) interface{})
		}
		return nil
	}
}

var ByItself = generic.DefineFunc("MaxByItself(vals T) E").
	Param("T", "array type").
	Param("E", "array element type", func(argMap generic.ArgMap) interface{} {
	return argMap["T"].(reflect.Type).Elem()
}).
	ImportFunc(compare.ByItself).
	Source(`
{{ $compare := expand "CompareByItself" "T" .E }}
currentMax := vals[0]
for i := 1; i < len(vals); i++ {
	if {{$compare}}(vals[i], currentMax) > 0 {
		currentMax = vals[i]
	}
}
return currentMax`)

var ByItselfForPlz = generic.DefineFunc("MaxByItselfForPlz(vals []interface{}) interface{}").
	Param("T", "array element type").
	ImportFunc(compare.ByItself).
	Source(`
{{ $compare := expand "CompareByItself" "T" .T }}
currentMax := vals[0].({{.T|name}})
for i := 1; i < len(vals); i++ {
	typedVal := vals[i].({{.T|name}})
	if {{$compare}}(typedVal, currentMax) > 0 {
		currentMax = typedVal
	}
}
return currentMax`)
