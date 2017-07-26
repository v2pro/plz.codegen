package cp2

import "github.com/v2pro/wombat/generic"

func init() {
	Anything.ImportFunc(copySimpleValue)
}

var copySimpleValue = generic.DefineFunc("CopySimpleValue(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	Source(`
if dst != nil {
	*dst = ({{.DT|elem|name}})(src)
}
`)