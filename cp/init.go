package cp

import (
	// register functions into cpAnything
	_ "github.com/v2pro/wombat/cp/cpArrayToArray"
	_ "github.com/v2pro/wombat/cp/cpFromInterface"
	_ "github.com/v2pro/wombat/cp/cpFromPtr"
	_ "github.com/v2pro/wombat/cp/cpIntoInterface"
	_ "github.com/v2pro/wombat/cp/cpIntoPtr"
	_ "github.com/v2pro/wombat/cp/cpMapToMap"
	_ "github.com/v2pro/wombat/cp/cpMapToStruct"
	_ "github.com/v2pro/wombat/cp/cpSimpleValue"
	_ "github.com/v2pro/wombat/cp/cpSliceToSlice"
	"github.com/v2pro/wombat/cp/cpAnything"
	_ "github.com/v2pro/wombat/cp/cpAnything"
	_ "github.com/v2pro/wombat/cp/cpStructToMap"
	_ "github.com/v2pro/wombat/cp/cpStructToStruct"
	"reflect"
)

// Gen generates a instance of cpAnything
func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	return cpAnything.Gen(dstType, srcType)
}
