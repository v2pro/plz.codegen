package compare

import "github.com/v2pro/wombat/generic"

func init() {
	ByItself.ImportFunc(comparePtr)
}

var comparePtr = generic.DefineFunc("ComparePtr(val1 T, val2 T) int").
	Param("T", "the type of value to compare").
	ImportFunc(ByItself).
	Source(`
{{ $compare := expand "CompareByItself" "T" (.T|elem) }}
return {{$compare}}(*val1, *val2)`)
