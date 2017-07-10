package fp_max

type funcMaxFloat64 struct {
}

func (f *funcMaxFloat64) max(collection []interface{}) interface{} {
	currentMax := collection[0].(float64)
	for _, elemObj := range collection[1:] {
		elem := elemObj.(float64)
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}
