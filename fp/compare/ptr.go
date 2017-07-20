package compare

import "github.com/v2pro/wombat/generic"

var comparePtr = generic.Func("ComparePtr").
	Params("T", "the type of value to compare").
	ImportFunc(F).
	Source(`
{{ $compare := expand "Compare" "T" (.T|elem) }}
func {{.funcName}}(val1 {{.T|name}}, val2 {{.T|name}}) int {
	return {{$compare}}(*val1, *val2)
}`)