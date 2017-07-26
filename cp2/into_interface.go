package cp2

import "github.com/v2pro/wombat/generic"

func init() {
	Anything.ImportFunc(copyIntoInterface)
}

var copyIntoInterface = generic.DefineFunc("CopyIntoInterface(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Source(`
if *dst == nil {
	newDst := new({{.ST|name}})
	newErr := copyDynamically(newDst, src)
	if newErr != nil && *err == nil {
		*err = newErr
	}
	*dst = *newDst
} else {
	newErr := copyDynamically(*dst, src)
	if newErr != nil && *err == nil {
		*err = newErr
	}
}`)
