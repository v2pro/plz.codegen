package cp

import (
	"github.com/v2pro/plz/acc"
)

func structToMap(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	fieldCopiers, err := createStructToMapFieldCopiers(dstAcc, srcAcc)
	if err != nil {
		return nil, err
	}
	return &structToMapCopier{
		fieldCopiers: fieldCopiers,
		dstAcc: dstAcc,
		dstKeyAcc: dstAcc.Key(),
	}, nil
}

func createStructToMapFieldCopiers(dstAcc acc.Accessor, srcAcc acc.Accessor) (map[string]Copier, error) {
	fieldCopiers := map[string]Copier{}
	dstElemAcc := dstAcc.Elem()
	for i := 0; i < srcAcc.NumField(); i++ {
		field := srcAcc.Field(i)
		copier, err := CopierOf(dstElemAcc, field.Accessor)
		if err != nil {
			return nil, err
		}
		fieldCopiers[field.Name] = copier
	}
	return fieldCopiers, nil
}

type structToMapCopier struct {
	fieldCopiers map[string]Copier
	dstAcc acc.Accessor
	dstKeyAcc acc.Accessor
}

func (copier *structToMapCopier) Copy(dst interface{}, src interface{}) (err error) {
	for fieldName, fieldCopier := range copier.fieldCopiers {
		copier.dstAcc.SetMap(dst, func(dstKey interface{}, dstElem interface{}) {
			copier.dstKeyAcc.SetString(dstKey, fieldName)
			err = fieldCopier.Copy(dstElem, src)
		})
	}
	return
}

