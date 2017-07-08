package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
	"github.com/v2pro/plz/logging"
	"reflect"
	"fmt"
)

func newStructToStructCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%#v => %#v] begin setup struct to struct copier", srcAcc, dstAcc))
	}
	fieldCopiers := make([]util.Copier, srcAcc.NumField())
	dstFields := map[string]lang.StructField{}
	for i := 0; i < dstAcc.NumField(); i++ {
		field := dstAcc.Field(i)
		dstFields[field.Name()] = field
	}
	for i := 0; i < srcAcc.NumField(); i++ {
		field := srcAcc.Field(i)
		dstField := dstFields[field.Name()]
		if dstField == nil {
			if shouldDebug {
				logger.Debug("skip src field", "fieldName", field.Name())
			}
			fieldCopiers[i] = &skipCopier{field.Name(), field.Accessor()}
		} else {
			if shouldDebug {
				logger.Debug("matched field", "fieldName", field.Name())
			}
			copier, err := util.CopierOf(dstField.Accessor(), field.Accessor())
			if err != nil {
				return nil, err
			}
			fieldCopiers[i] = &structFieldCopier{
				elemCopier: copier,
				dstIndex:   dstField.Index(),
				dstAcc:     dstAcc,
				srcAcc:     field.Accessor(),
			}
		}
	}
	if shouldDebug {
		srcFields := map[string]lang.StructField{}
		for i := 0; i < srcAcc.NumField(); i++ {
			field := srcAcc.Field(i)
			srcFields[field.Name()] = field
		}
		for i := 0; i < dstAcc.NumField(); i++ {
			field := dstAcc.Field(i)
			srcField := srcFields[field.Name()]
			if srcField == nil {
				logger.Debug("skip dst field", "fieldName", field.Name())
			}
		}
		logger.Debug(fmt.Sprintf("[%#v => %#v] end setup struct to struct copier", srcAcc, dstAcc))
	}
	return &structToStructCopier{
		fieldCopiers: fieldCopiers,
		srcAcc:       srcAcc,
		dstAcc:       dstAcc,
	}, nil
}

type structToStructCopier struct {
	fieldCopiers []util.Copier
	srcAcc       lang.Accessor
	dstAcc       lang.Accessor
}

func (copier *structToStructCopier) Copy(dst, src unsafe.Pointer) (err error) {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	if copier.srcAcc.IsNil(src) {
		if shouldDebug {
			logger.Debug("skip copy struct as src is nil", "src", src, "dst", dst,
				"srcAcc", reflect.TypeOf(copier.srcAcc),
				"dstAcc", reflect.TypeOf(copier.dstAcc))
		}
		copier.dstAcc.Skip(dst)
		return nil
	}
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] begin copy struct to struct", src), "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc))
	}
	copier.srcAcc.IterateArray(src, func(index int, srcElem unsafe.Pointer) bool {
		err = copier.fieldCopiers[index].Copy(dst, srcElem)
		if err != nil {
			return false
		}
		return true
	})
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] end copy struct to struct", src), "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc))
	}
	return
}

type structFieldCopier struct {
	elemCopier util.Copier
	dstIndex   int
	dstAcc     lang.Accessor
	srcAcc     lang.Accessor
}

func (copier *structFieldCopier) Copy(dst, src unsafe.Pointer) error {
	shouldDebug := logger.ShouldLog(logging.LEVEL_DEBUG)
	if copier.srcAcc.IsNil(src) {
		logger.Debug("skip copy struct field as src is nil", "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"fieldName", copier.dstAcc.Field(copier.dstIndex).Name())
		copier.dstAcc.Skip(dst)
		return nil
	}
	dstElem := copier.dstAcc.ArrayIndex(dst, copier.dstIndex)
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] begin copy struct field to struct field", src), "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"fieldName", copier.dstAcc.Field(copier.dstIndex).Name(),
			"dstElem", dstElem,
		"elemCopier", reflect.TypeOf(copier.elemCopier))
	}
	err := copier.elemCopier.Copy(dstElem, src)
	if shouldDebug {
		logger.Debug(fmt.Sprintf("[%x] end copy struct field to struct field", src), "src", src, "dst", dst,
			"srcAcc", reflect.TypeOf(copier.srcAcc),
			"dstAcc", reflect.TypeOf(copier.dstAcc),
			"fieldName", copier.dstAcc.Field(copier.dstIndex).Name(),
			"dstElem", dstElem,
			"elemCopier", reflect.TypeOf(copier.elemCopier),
			"err", err)
	}
	return err
}
