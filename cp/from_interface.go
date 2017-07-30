package cp

import "github.com/v2pro/wombat/generic"

func init() {
	Anything.ImportFunc(copyFromInterface)
}

var copyFromInterface = generic.DefineFunc("CopyFromInterface(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Source(`
newErr := copyDynamically(dst, src)
if newErr != nil && *err == nil {
	*err = newErr
}`)
