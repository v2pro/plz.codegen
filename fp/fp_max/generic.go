package fp_max

import (
	"github.com/v2pro/plz/lang"
)

type funcMaxGeneric struct {
	comparator lang.Comparator
}

func (f *funcMaxGeneric) max(collection []interface{}) interface{} {
	currentMax := collection[0]
	for _, elemObj := range collection[1:] {
		if objType(elemObj) != objType(currentMax) {
			panic("type mismatch")
		}
		if f.comparator.Compare(objPtr(elemObj), objPtr(currentMax)) > 0 {
			currentMax = elemObj
		}
	}
	return currentMax
}
