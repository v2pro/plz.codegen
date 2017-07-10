package fp_max

type funcMaxInt struct {
}

func (f *funcMaxInt) max(collection []interface{}) interface{} {
	currentMax := collection[0].(int)
	for _, elemObj := range collection[1:] {
		elem := elemObj.(int)
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}
