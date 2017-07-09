package fp_max

import (
	"reflect"
	"fmt"
)

type maxInt struct {
}

const maxUint = ^uint(0)
const minInt = -int(maxUint >> 1) - 1

func (f *maxInt) max(collection []interface{}) interface{} {
	currentMax := minInt
	for _, elemObj := range collection {
		if objKind(elemObj) != reflect.Int {
			panic(fmt.Sprintf("type mismatch, expect %v but found %v",
				reflect.Int, reflect.TypeOf(elemObj)))
		}
		elem := *(*int)(objPtr(elemObj))
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}
