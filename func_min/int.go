package func_max

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"fmt"
)

type minInt struct {
	accessor lang.Accessor
}

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func (f *minInt) min(collection []interface{}) interface{} {
	currentMin := maxInt
	for _, elemObj := range collection {
		kind := reflect.TypeOf(elemObj).Kind()
		if kind != reflect.Int {
			panic(fmt.Sprintf("type mismatch, expect %v but found %v", reflect.Int, kind))
		}
		elem := f.accessor.Int(lang.AddressOf(elemObj))
		if elem < currentMin {
			currentMin = elem
		}
	}
	return currentMin
}
