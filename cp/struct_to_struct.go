package cp

import "github.com/v2pro/plz/acc"

func structToStruct(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	fieldCopiers, err := createFieldCopiers(dstAcc, srcAcc)
	if err != nil {
		return nil, err
	}
	return &structToStructCopier{fieldCopiers}, nil
}

func createFieldCopiers(dstAcc acc.Accessor, srcAcc acc.Accessor) (map[string]Copier, error) {
	bindings := map[string]*binding{}
	for i := 0; i < dstAcc.NumField(); i++ {
		field := dstAcc.Field(i)
		bindings[field.Name] = &binding{
			dstAcc: field.Accessor,
		}
	}
	for i := 0; i < srcAcc.NumField(); i++ {
		field := srcAcc.Field(i)
		binding := bindings[field.Name]
		if binding == nil {
			continue
		}
		binding.srcAcc = field.Accessor
	}
	fieldCopiers := map[string]Copier{}
	for fieldName, v := range bindings {
		if v.srcAcc != nil && v.dstAcc != nil {
			copier, err := CopierOf(v.dstAcc, v.srcAcc)
			if err != nil {
				return nil, err
			}
			fieldCopiers[fieldName] = copier
		}
	}
	return fieldCopiers, nil
}

type binding struct {
	srcAcc acc.Accessor
	dstAcc acc.Accessor
}

type structToStructCopier struct {
	fieldCopiers map[string]Copier
}

func (copier *structToStructCopier) Copy(dst interface{}, src interface{}) error {
	for _, copier := range copier.fieldCopiers {
		err := copier.Copy(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}
