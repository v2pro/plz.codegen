package cp

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"github.com/v2pro/plz/logging"
)

type skipCopier struct {
	fieldName string
	srcAcc lang.Accessor
}

func (copier *skipCopier) Copy(dst, src unsafe.Pointer) error {
	if logger.ShouldLog(logging.LEVEL_DEBUG) {
		logger.Debug("skip struct field as field name not matched",
		"src", src, "dst", dst,
		"fieldName", copier.fieldName)
	}
	copier.srcAcc.Skip(src)
	return nil
}
