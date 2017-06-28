package wombat

import (
	"github.com/v2pro/plz"
	"reflect"
	"fmt"
	_ "github.com/v2pro/plz/acc/native"
	"github.com/v2pro/plz/acc"
)

func Copy(dst interface{}, src interface{}) error {
	dstAcc := plz.AccessorOf(reflect.TypeOf(dst))
	srcAcc := plz.AccessorOf(reflect.TypeOf(src))
	copier, err := copierOf(dstAcc, srcAcc)
	if err != nil {
		return err
	}
	return copier.Copy(dst, src)
}

func copierOf(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
	if dstAcc.Kind() != srcAcc.Kind() {
		return nil, fmt.Errorf("kind mismatch: %v", dstAcc, srcAcc)
	}
	switch dstAcc.Kind() {
	case reflect.Struct:
		return copierOfStruct(dstAcc, srcAcc)
	case reflect.String:
		return &stringCopier{dstAcc: dstAcc, srcAcc: srcAcc}, nil
	default:
		return nil, fmt.Errorf("do not know how to copy %#v => %#v", srcAcc, dstAcc)
	}
}

func copierOfStruct(dstAcc acc.Accessor, srcAcc acc.Accessor) (Copier, error) {
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
	fieldCopiers := []Copier{}
	for _, v := range bindings {
		if v.srcAcc != nil && v.dstAcc != nil {
			copier, err := copierOf(v.dstAcc, v.srcAcc)
			if err != nil {
				return nil, err
			}
			fieldCopiers = append(fieldCopiers, copier)
		}
	}
	return &structCopier{fieldCopiers}, nil
}

type binding struct {
	srcAcc acc.Accessor
	dstAcc acc.Accessor
}

type Copier interface {
	Copy(dst interface{}, src interface{}) error
}

type structCopier struct {
	fieldCopiers []Copier
}

func (copier *structCopier) Copy(dst interface{}, src interface{}) error {
	for _, copier := range copier.fieldCopiers {
		err := copier.Copy(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}

type stringCopier struct {
	srcAcc acc.Accessor
	dstAcc acc.Accessor
}

func (copier stringCopier) Copy(dst interface{}, src interface{}) error {
	copier.dstAcc.SetString(dst, copier.srcAcc.String(src))
	return nil
}
