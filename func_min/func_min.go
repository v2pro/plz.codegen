package func_max

import (
	"github.com/v2pro/plz/util"
	"github.com/v2pro/plz/lang"
	"reflect"
	_ "github.com/v2pro/wombat/acc"
	"fmt"
)

func init() {
	util.Min = min
}

func min(collection ...interface{}) interface{} {
	if len(collection) == 0 {
		return nil
	}
	typ := reflect.TypeOf(collection[0])
	f := funcMinMap[typ]
	if f == nil {
		panic(fmt.Sprintf("no min implementation for: %v", typ))
	}
	return f.min(collection)
}

var funcMinMap = map[reflect.Type]funcMin{
	reflect.TypeOf(0): &minInt{accessor:lang.AccessorOf(reflect.TypeOf(0), "")},
}

type funcMin interface {
	min(collection []interface{}) interface{}
}

