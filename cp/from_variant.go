package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
	"github.com/v2pro/plz/logging"
	"fmt"
	"reflect"
)

type fromVariantCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *fromVariantCopier) Copy(dst, src unsafe.Pointer) error {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	if copier.srcAcc.IsNil(src) {
		if shouldDebug {
			logger.Debug("skip copy from variant as src is nil", "src", src, "dst", dst)
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	srcElem, srcElemAcc := copier.srcAcc.VariantElem(src)
	if srcElem == nil {
		if shouldDebug {
			logger.Debug(
				"skip copy from variant as variant unwrapped is nil",
				"src", src, "dst", dst,
				"srcAcc", reflect.TypeOf(copier.srcAcc),
				"dstAcc", reflect.TypeOf(copier.dstAcc))
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	elemCopier, err := util.CopierOf(copier.dstAcc, srcElemAcc)
	if err != nil {
		return err
	}
	if shouldDebug {
		logger.Debug(
			fmt.Sprintf("[%x] begin copy variant", srcElem),
			"src", src, "dst", dst,
			"srcElem", srcElem, "srcElemAcc", reflect.TypeOf(srcElemAcc),
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"elemCopier", reflect.TypeOf(elemCopier))
	}
	err = elemCopier.Copy(dst, srcElem)
	if shouldDebug {
		logger.Debug(
			fmt.Sprintf("[%x] end copy variant", srcElem),
			"src", src, "dst", dst,
			"srcElem", srcElem, "srcElemAcc", reflect.TypeOf(srcElemAcc), "err", err,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc))
	}
	return err
}
