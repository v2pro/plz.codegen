package fp_max

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"fmt"
)

type maxInt struct {
	accessor lang.Accessor
}

const maxUint = ^uint(0)
const minInt = -int(maxUint >> 1) - 1

func (f *maxInt) max(collection []interface{}) interface{} {
	currentMax := minInt
	for _, elemObj := range collection {
		kind := reflect.TypeOf(elemObj).Kind()
		if kind != reflect.Int {
			panic(fmt.Sprintf("type mismatch, expect %v but found %v", reflect.Int, kind))
		}
		elem := f.accessor.Int(lang.AddressOf(elemObj))
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}
