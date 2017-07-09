package func_max

import (
	"github.com/v2pro/plz/util"
	"github.com/v2pro/plz/lang"
	"reflect"
	_ "github.com/v2pro/wombat/acc"
	"fmt"
)

func init() {
	util.Max = max
}

func max(collection ...interface{}) interface{} {
	if len(collection) == 0 {
		return nil
	}
	typ := reflect.TypeOf(collection[0])
	f := funcMaxMap[typ]
	if f == nil {
		panic(fmt.Sprintf("no max implementation for: %v", typ))
	}
	return f.max(collection)
}

var funcMaxMap = map[reflect.Type]funcMax {
	reflect.TypeOf(0): &maxInt{accessor:lang.AccessorOf(reflect.TypeOf(0), "")},
}

type funcMax interface {
	max(collection []interface{}) interface{}
}

