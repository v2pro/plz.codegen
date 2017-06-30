package cp

import "github.com/v2pro/plz/acc"

func mapToStruct(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	fieldCopiers, err := createFieldCopiers(dstAcc, srcAcc)
	if err != nil {
		return nil, err
	}
	return &mapToStructCopier{
		fieldCopiers: fieldCopiers,
		srcAcc: srcAcc,
		srcKeyAcc: srcAcc.Key(),
		srcElemAcc: srcAcc.Elem(),
	}, nil
}

type mapToStructCopier struct {
	fieldCopiers map[string]Copier
	srcAcc       acc.Accessor
	srcKeyAcc    acc.Accessor
	srcElemAcc   acc.Accessor
}

func (copier *mapToStructCopier) Copy(dst interface{}, src interface{}) error {
	copier.srcAcc.IterateMap(src, func(key interface{}, elem interface{}) bool {
		fieldName := copier.srcKeyAcc.String(key)
		fieldCopier := copier.fieldCopiers[fieldName]
		if fieldCopier == nil {
			copier.srcElemAcc.Skip(elem)
		} else {
			fieldCopier.Copy(dst, elem)
		}
		return true
	})
	return nil
}
