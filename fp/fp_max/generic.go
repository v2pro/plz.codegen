package fp_max

import (
	"github.com/v2pro/plz/lang"
)

type funcMaxGeneric struct {
	comparator lang.ObjectComparator
}

func (f *funcMaxGeneric) max(collection []interface{}) interface{} {
	currentMax := collection[0]
	for _, elemObj := range collection[1:] {
		if f.comparator.Compare(elemObj, currentMax) > 0 {
			currentMax = elemObj
		}
	}
	return currentMax
}
