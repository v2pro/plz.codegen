package cp

import (
	_ "github.com/v2pro/wombat/cp/cpStatically"
	_ "github.com/v2pro/wombat/cp/cpIntoPtr"
	_ "github.com/v2pro/wombat/cp/cpFromPtr"
	_ "github.com/v2pro/wombat/cp/cpSimpleValue"
	_ "github.com/v2pro/wombat/cp/cpStructToStruct"
	_ "github.com/v2pro/wombat/cp/cpStructToMap"
	_ "github.com/v2pro/wombat/cp/cpMapToMap"
	_ "github.com/v2pro/wombat/cp/cpMapToStruct"
	_ "github.com/v2pro/wombat/cp/cpSliceToSlice"
	_ "github.com/v2pro/wombat/cp/cpArrayToArray"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
)

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	return cpStatically.Gen(dstType, srcType)
}
