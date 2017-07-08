package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/v2pro/plz/logging"
)

type intCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *intCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	if copier.srcAcc.IsNil(src) {
		if logger.ShouldLog(logging.LEVEL_DEBUG) {
			logger.Debug("skip copy int as src is nil", "src", src, "dst", dst)
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	val := copier.srcAcc.Int(src)
	if logger.ShouldLog(logging.LEVEL_DEBUG) {
		logger.Debug("copy int", "val", val, "src", src, "dst", dst)
	}
	copier.dstAcc.SetInt(dst, val)
	return nil
}
