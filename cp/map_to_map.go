package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
	"github.com/v2pro/plz/logging"
	"reflect"
	"fmt"
)

func newMapToMapCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
	keyCopier, err := util.CopierOf(dstAcc.Key(), srcAcc.Key())
	if err != nil {
		return nil, err
	}
	elemCopier, err := util.CopierOf(dstAcc.Elem(), srcAcc.Elem())
	if err != nil {
		return nil, err
	}
	return &mapToMapCopier{
		srcAcc:     srcAcc,
		dstAcc:     dstAcc,
		keyCopier:  keyCopier,
		elemCopier: elemCopier,
	}, nil
}

type mapToMapCopier struct {
	srcAcc     lang.Accessor
	dstAcc     lang.Accessor
	keyCopier  util.Copier
	elemCopier util.Copier
}

func (copier *mapToMapCopier) Copy(dst, src unsafe.Pointer) (err error) {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	if copier.srcAcc.IsNil(src) {
		if shouldDebug {
			logger.Debug("skip copy map to map", "src", src, "dst", dst,
				"srcAcc", reflect.TypeOf(copier.srcAcc),
				"dstAcc", reflect.TypeOf(copier.dstAcc),
				"keyCopier", reflect.TypeOf(copier.keyCopier),
				"elemCopier", reflect.TypeOf(copier.elemCopier))
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] begin copy map to map", src),
			"src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"keyCopier", reflect.TypeOf(copier.keyCopier),
			"elemCopier", reflect.TypeOf(copier.elemCopier))
	}
	copier.dstAcc.FillMap(dst, func(filler lang.MapFiller) {
		copier.srcAcc.IterateMap(src, func(srcKey unsafe.Pointer, srcElem unsafe.Pointer) bool {
			if copier.dstAcc.RandomAccessible() {
				dstElem := copier.dstAcc.MapIndex(dst, srcKey)
				if dstElem != nil {
					if shouldDebug {
						logger.Debug(fmt.Sprintf("[%x] begin copy map entry to existing map entry", srcKey),
							"srcKey", srcKey, "src", src, "dst", dst,
							"srcAcc", reflect.TypeOf(copier.srcAcc),
							"dstAcc", reflect.TypeOf(copier.dstAcc),
							"keyCopier", reflect.TypeOf(copier.keyCopier),
							"elemCopier", reflect.TypeOf(copier.elemCopier))
					}
					err = copier.elemCopier.Copy(dstElem, srcElem)
					if shouldDebug {
						logger.Debug(fmt.Sprintf("[%x] end copy map entry to existing map entry", srcKey),
							"srcKey", srcKey, "src", src, "dst", dst,
							"srcAcc", reflect.TypeOf(copier.srcAcc),
							"dstAcc", reflect.TypeOf(copier.dstAcc),
							"keyCopier", reflect.TypeOf(copier.keyCopier),
							"elemCopier", reflect.TypeOf(copier.elemCopier),
							"err", err)
					}
					if err != nil {
						return false
					}
					copier.dstAcc.SetMapIndex(dst, srcKey, dstElem)
					return true
				}
			}
			if shouldDebug {
				logger.Debug(fmt.Sprintf("[%x] begin copy map entry to new map entry", srcKey),
					"srcKey", srcKey, "src", src, "dst", dst,
					"srcAcc", reflect.TypeOf(copier.srcAcc),
					"dstAcc", reflect.TypeOf(copier.dstAcc),
					"keyCopier", reflect.TypeOf(copier.keyCopier),
					"elemCopier", reflect.TypeOf(copier.elemCopier))
			}
			dstKey, dstElem := filler.Next()
			err = copier.keyCopier.Copy(dstKey, srcKey)
			if err != nil {
				return false
			}
			err = copier.elemCopier.Copy(dstElem, srcElem)
			if shouldDebug {
				logger.Debug(fmt.Sprintf("[%x] end copy map entry to new map entry", srcKey),
					"srcKey", srcKey, "src", src, "dst", dst,
					"srcAcc", reflect.TypeOf(copier.srcAcc),
					"dstAcc", reflect.TypeOf(copier.dstAcc),
					"keyCopier", reflect.TypeOf(copier.keyCopier),
					"elemCopier", reflect.TypeOf(copier.elemCopier),
					"err", err)
			}
			if err != nil {
				return false
			}
			filler.Fill()
			return true
		})
	})
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] end copy map to map", src),
			"src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"keyCopier", reflect.TypeOf(copier.keyCopier),
			"elemCopier", reflect.TypeOf(copier.elemCopier),
			"err", err)
	}
	return
}
