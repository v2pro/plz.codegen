package wombat

import (
	"unsafe"
	"github.com/v2pro/plz/util"
	"github.com/v2pro/plz/lang"
	"fmt"
)

func newStructToStructCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
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
			fieldCopiers[i] = &skipCopier{field.Accessor()}
		} else {
			copier, err := util.CopierOf(dstField.Accessor(), field.Accessor())
			if err != nil {
				return nil, err
			}
			if copier == nil {
				return nil, fmt.Errorf("no copier for %#v => %#v", field.Accessor(), dstField.Accessor())
			}
			fieldCopiers[i] = &structFieldCopier{
				elemCopier: copier,
				dstIndex:   dstField.Index(),
				dstAcc:     dstAcc,
			}
		}
	}
	return &structToStructCopier{
		fieldCopiers: fieldCopiers,
		srcAcc:       srcAcc,
	}, nil
}

type structToStructCopier struct {
	fieldCopiers []util.Copier
	srcAcc       lang.Accessor
}

func (copier *structToStructCopier) Copy(dst, src unsafe.Pointer) (err error) {
	copier.srcAcc.IterateArray(src, func(index int, srcElem unsafe.Pointer) bool {
		err = copier.fieldCopiers[index].Copy(dst, srcElem)
		if err != nil {
			return false
		}
		return true
	})
	return
}

type structFieldCopier struct {
	elemCopier util.Copier
	dstIndex   int
	dstAcc     lang.Accessor
}

func (copier *structFieldCopier) Copy(dst, src unsafe.Pointer) error {
	dstElem := copier.dstAcc.ArrayIndex(dst, copier.dstIndex)
	return copier.elemCopier.Copy(dstElem, src)
}
