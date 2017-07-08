package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/v2pro/plz/logging"
)

type stringCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *stringCopier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	if copier.srcAcc.IsNil(src) {
		if logger.ShouldLog(logging.LEVEL_DEBUG) {
			logger.Debug("skip copy string as src is nil", "src", src, "dst", dst)
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	val := copier.srcAcc.String(src)
	if logger.ShouldLog(logging.LEVEL_DEBUG) {
		logger.Debug("copy string", "val", val, "src", src, "dst", dst)
	}
	copier.dstAcc.SetString(dst, val)
	return nil
}
