package wombat

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"unsafe"
)

func newMapToStructCopier(dstAcc, srcAcc lang.Accessor) (util.Copier, error) {
	bindings := map[string]*binding{}
	for i := 0; i < dstAcc.NumField(); i++ {
		field := dstAcc.Field(i)
		copier, err := util.CopierOf(field.Accessor(), srcAcc.Elem())
		if err != nil {
			return nil, err
		}
		bindings[field.Name()] = &binding{
			index:  i,
			copier: copier,
		}
	}
	return &mapToStructCopier{
		srcKeyAcc:  srcAcc.Key(),
		srcElemAcc: srcAcc.Elem(),
		srcAcc:     srcAcc,
		dstAcc:     dstAcc,
		bindings:   bindings,
	}, nil
}

type mapToStructCopier struct {
	srcKeyAcc  lang.Accessor
	srcElemAcc lang.Accessor
	srcAcc     lang.Accessor
	dstAcc     lang.Accessor
	bindings   map[string]*binding
}

type binding struct {
	index  int
	copier util.Copier
}

func (copier *mapToStructCopier) Copy(dst, src unsafe.Pointer) (err error) {
	copier.srcAcc.IterateMap(src, func(srcKey unsafe.Pointer, srcElem unsafe.Pointer) bool {
		fieldName := copier.srcKeyAcc.String(srcKey)
		binding := copier.bindings[fieldName]
		if binding == nil {
			copier.srcElemAcc.Skip(srcElem)
		} else {
			dstElem := copier.dstAcc.ArrayIndex(dst, binding.index)
			err = binding.copier.Copy(dstElem, srcElem)
			if err != nil {
				return false
			}
		}
		return true
	})
	return
}
