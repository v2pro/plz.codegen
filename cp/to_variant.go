package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
	"github.com/v2pro/plz/logging"
	"reflect"
	"fmt"
)

type toVariantCopier struct {
	srcAcc lang.Accessor
	dstAcc lang.Accessor
}

func (copier *toVariantCopier) Copy(dst, src unsafe.Pointer) error {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	dstElem, dstElemAcc := copier.dstAcc.VariantElem(dst)
	if dstElem == nil {
		if shouldDebug {
			logger.Debug("dst variant is nil, initialize", "src", src, "dst", dst,
				"srcAcc", reflect.TypeOf(copier.srcAcc),
				"dstAcc", reflect.TypeOf(copier.dstAcc))
		}
		dstElem, dstElemAcc = copier.dstAcc.InitVariant(dst, copier.srcAcc)
		if shouldDebug {
			logger.Debug("dst variant initialized", "src", src, "dst", dst,
				"srcAcc", reflect.TypeOf(copier.srcAcc),
				"dstAcc", reflect.TypeOf(copier.dstAcc),
				"dstElem", dstElem,
				"dstElemAcc", reflect.TypeOf(dstElemAcc))
		}
	}
	elemCopier, err := util.CopierOf(dstElemAcc, copier.srcAcc)
	if err != nil {
		return err
	}
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] begin copy to variant", src), "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"dstElem", dstElem,
			"dstElemAcc", reflect.TypeOf(dstElemAcc),
			"elemCopier", reflect.TypeOf(elemCopier))
	}
	err = elemCopier.Copy(dstElem, src)
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] end copy to variant", src), "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"dstElem", dstElem,
			"dstElemAcc", reflect.TypeOf(dstElemAcc),
			"elemCopier", reflect.TypeOf(elemCopier),
			"err", err)
	}
	return err
}
