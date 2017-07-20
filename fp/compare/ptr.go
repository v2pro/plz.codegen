package compare

import "github.com/v2pro/wombat/generic"

func init() {
	ByItself.ImportFunc(comparePtr)
}

var comparePtr = generic.Func("ComparePtr(val1 T, val2 T) int").
	Params("T", "the type of value to compare").
	ImportFunc(ByItself).
	Source(`
{{ $compare := expand "Compare" "T" (.T|elem) }}
return {{$compare}}(*val1, *val2)`)
