package compare

import "github.com/v2pro/wombat/generic"

var compareSimpleValue = generic.Func("CompareSimpleValue").
	Params("T", "the type of value to compare").
	Source(`
	func {{.funcName}}(val1 {{.T|name}}, val2 {{.T|name}}) int {
		if val1 < val2 {
			return -1
		} else if val1 == val2 {
			return 0
		} else {
			return 1
		}
	}`)
