package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/v2pro/plz/logging"
)

type float64Copier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *float64Copier) Copy(dst unsafe.Pointer, src unsafe.Pointer) error {
	if copier.srcAcc.IsNil(src) {
		if logger.ShouldLog(logging.LEVEL_DEBUG) {
			logger.Debug("skip copy float64 as src is nil", "src", src, "dst", dst)
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	val := copier.srcAcc.Float64(src)
	if logger.ShouldLog(logging.LEVEL_DEBUG) {
		logger.Debug("copy float64", "val", val, "src", src, "dst", dst)
	}
	copier.dstAcc.SetFloat64(dst, val)
	return nil
}
