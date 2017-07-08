package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
	"github.com/v2pro/plz/logging"
	"fmt"
)

func newStructToMapCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
	fieldCopiers := make([]util.Copier, srcAcc.NumField())
	fieldAccessors := make([]lang.Accessor, srcAcc.NumField())
	for i := 0; i < srcAcc.NumField(); i++ {
		field := srcAcc.Field(i)
		copier, err := util.CopierOf(dstAcc.Elem(), field.Accessor())
		if err != nil {
			return nil, err
		}
		fieldAccessors[i] = field.Accessor()
		fieldCopiers[i] = copier
	}
	return &structToMapCopier{
		fieldCopiers: fieldCopiers,
		fieldAccessors: fieldAccessors,
		srcAcc:       srcAcc,
		dstAcc:       dstAcc,
	}, nil
}

type structToMapCopier struct {
	fieldCopiers []util.Copier
	fieldAccessors []lang.Accessor
	srcAcc       lang.Accessor
	dstAcc       lang.Accessor
}

func (copier *structToMapCopier) Copy(dst, src unsafe.Pointer) (err error) {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] begin copy struct to map", src), "src", src, "dst", dst)
	}
	copier.dstAcc.FillMap(dst, func(filler lang.MapFiller) {
		copier.srcAcc.IterateArray(src, func(index int, srcElem unsafe.Pointer) bool {
			field := copier.srcAcc.Field(index)
			fieldName := field.Name()
			if copier.fieldAccessors[index].IsNil(srcElem) {
				if shouldDebug {
					logger.Debug("skip copy struct field to map, as field is nil",
						"fieldName", fieldName, "src", src, "dst", dst, "srcElem", srcElem)
				}
				// TODO: copy into existing map entry as nil value
				dstKey, dstElem := filler.Next()
				copier.dstAcc.Key().SetString(dstKey, fieldName)
				copier.dstAcc.Elem().Skip(dstElem)
				filler.Fill()
				return true
			}
			if copier.dstAcc.RandomAccessible() {
				dstElem := copier.dstAcc.MapIndex(dst, lang.AddressOf(fieldName))
				if dstElem != nil {
					if shouldDebug {
						logger.Debug(fmt.Sprintf("[%x] begin copy struct field to existing map entry",
							srcElem), "fieldName", fieldName, "src", src, "dst", dst)
					}
					err = copier.fieldCopiers[index].Copy(dstElem, srcElem)
					if shouldDebug {
						logger.Debug(fmt.Sprintf("[%x] end copy struct field to existing map entry", srcElem),
							"fieldName", fieldName, "src", src, "dst", dst)
					}
					if err != nil {
						return false
					}
					copier.dstAcc.SetMapIndex(dst, lang.AddressOf(fieldName), dstElem)
					return true
				}
			}
			if shouldDebug {
				logger.Debug(fmt.Sprintf("[%x] begin copy struct field to new map entry", srcElem),
					"fieldName", fieldName, "src", src, "dst", dst)
			}
			dstKey, dstElem := filler.Next()
			copier.dstAcc.Key().SetString(dstKey, fieldName)
			err = copier.fieldCopiers[index].Copy(dstElem, srcElem)
			if shouldDebug {
				logger.Debug(fmt.Sprintf("[%x] end copy struct field to new map entry", srcElem),
					"fieldName", fieldName, "src", src, "dst", dst)
			}
			if err != nil {
				return false
			}
			filler.Fill()
			return true
		})
	})
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] end copy struct to map", src), "src", src, "dst", dst)
	}
	return
}
