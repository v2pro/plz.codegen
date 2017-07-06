package wombat

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
)

func newStructToMapCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
	fieldCopiers := make([]util.Copier, srcAcc.NumField())
	for i := 0; i < srcAcc.NumField(); i++ {
		field := srcAcc.Field(i)
		copier, err := util.CopierOf(dstAcc.Elem(), field.Accessor())
		if err != nil {
			return nil, err
		}
		fieldCopiers[i] = copier
	}
	return &structToMapCopier{
		fieldCopiers: fieldCopiers,
		srcAcc:       srcAcc,
		dstAcc:       dstAcc,
	}, nil
}

type structToMapCopier struct {
	fields       []lang.StructField
	fieldCopiers []util.Copier
	srcAcc       lang.Accessor
	dstAcc       lang.Accessor
}

func (copier *structToMapCopier) Copy(dst, src unsafe.Pointer) (err error) {
	copier.dstAcc.FillMap(dst, func(filler lang.MapFiller) {
		copier.srcAcc.IterateArray(src, func(index int, srcElem unsafe.Pointer) bool {
			dstKey, dstElem := filler.Next()
			err = copier.fieldCopiers[index].Copy(dstElem, srcElem)
			if err != nil {
				return false
			}
			copier.dstAcc.Key().SetString(dstKey, copier.srcAcc.Field(index).Name())
			filler.Fill()
			return true
		})
	})
	return
}
