package max

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/compare"
	"reflect"
)

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
