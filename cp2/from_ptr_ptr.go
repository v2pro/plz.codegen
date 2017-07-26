package cp2

import "github.com/v2pro/wombat/generic"

func init() {
	Anything.ImportFunc(copyFromPtrPtr)
}

var copyFromPtrPtr = generic.DefineFunc("CopyFromPtrPtr(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Source(`
{{ $cp := expand "CopyAnything" "DT" .DT "ST" (.ST|elem) }}
if src == nil {
	{{$cp}}(err, dst, nil)
	return
}
{{$cp}}(err, dst, *src)`)
