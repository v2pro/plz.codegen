package compare

import "github.com/v2pro/wombat/generic"

func init() {
	ByItself.ImportFunc(compareSimpleValue)
}

var compareSimpleValue = generic.Func("CompareSimpleValue(val1 T, val2 T) int").
	Params("T", "the type of value to compare").
	Source(`
if val1 < val2 {
	return -1
} else if val1 == val2 {
	return 0
} else {
	return 1
}`)
