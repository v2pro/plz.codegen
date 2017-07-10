package fp_max

import "github.com/v2pro/plz/lang"

type funcMaxComparable struct {
}

func (f *funcMaxComparable) max(collection []interface{}) interface{} {
	currentMax := collection[0].(lang.ObjectComparable)
	for _, elemObj := range collection[1:] {
		if currentMax.Compare(elemObj) < 0 {
			currentMax = elemObj.(lang.ObjectComparable)
		}
	}
	return currentMax
}
