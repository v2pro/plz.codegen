package cp

import "github.com/v2pro/wombat/generic"

func init() {
	Anything.ImportFunc(copyFromPtr)
}

var copyFromPtr = generic.DefineFunc("CopyFromPtr(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Source(`
{{ $cp := expand "CopyAnything" "DT" .DT "ST" (.ST|elem) }}
if src == nil {
	return
}
{{$cp}}(err, dst, *src)`)
