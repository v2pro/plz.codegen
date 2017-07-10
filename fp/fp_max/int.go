package fp_max

type funcMaxInt struct {
}

const maxUint = ^uint(0)
const minInt = -int(maxUint >> 1) - 1

func (f *funcMaxInt) max(collection []interface{}) interface{} {
	currentMax := minInt
	for _, elemObj := range collection {
		elem := elemObj.(int)
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}
