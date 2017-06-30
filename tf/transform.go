package tf

import (
	"github.com/v2pro/plz/acc"
)

func Transform(fromAccessor acc.Accessor, toAccessor acc.Accessor) (acc.Accessor, error) {
	if fromAccessor.Kind() == toAccessor.Kind() {
		return fromAccessor, nil
	}
	if fromAccessor.Kind() == acc.Interface {
		if toAccessor.Kind().IsSingleValue() {
			return fromAccessor, nil
		}
	}
	kind := toAccessor.Kind()
	if fromAccessor.Kind() != acc.Struct && kind == acc.Struct {
		kind = acc.Map
	}
	keyAccessor, err := Transform(fromAccessor.Key(), toAccessor.Key())
	if err != nil {
		return nil, err
	}
	elemAccessor, err := Transform(fromAccessor.Elem(), toAccessor.Elem())
	if err != nil {
		return nil, err
	}
	fields := []acc.StructField{}
	for i := 0; i < toAccessor.NumField(); i++ {
		field := toAccessor.Field(i)
		field.Accessor, err = Transform(fromAccessor.Elem(), field.Accessor)
		fields = append(fields, field)
		if err != nil {
			return nil, err
		}
	}
	return &transformedAccessor{
		kind:         kind,
		fields:       fields,
		keyAccessor:  keyAccessor,
		elemAccessor: elemAccessor,
		toAccessor:   toAccessor,
		fromAccessor: fromAccessor,
	}, nil
}

type transformedAccessor struct {
	acc.NoopAccessor
	kind         acc.Kind
	fields       []acc.StructField
	keyAccessor  acc.Accessor
	elemAccessor acc.Accessor
	toAccessor   acc.Accessor
	fromAccessor acc.Accessor
}

func (accessor *transformedAccessor) Kind() acc.Kind {
	return accessor.kind
}

func (accessor *transformedAccessor) GoString() string {
	return accessor.fromAccessor.GoString()
}

func (accessor *transformedAccessor) Key() acc.Accessor {
	return accessor.keyAccessor
}

func (accessor *transformedAccessor) Elem() acc.Accessor {
	return accessor.elemAccessor
}

func (accessor *transformedAccessor) NumField() int {
	return len(accessor.fields)
}

func (accessor *transformedAccessor) Field(index int) acc.StructField {
	return accessor.fields[index]
}

func (accessor *transformedAccessor) IterateMap(obj interface{}, cb func(key interface{}, elem interface{}) bool) {
	accessor.fromAccessor.IterateMap(obj, cb)
}

func (accessor *transformedAccessor) IterateArray(obj interface{}, cb func(elem interface{}) bool) {
	accessor.fromAccessor.IterateArray(obj, cb)
}

func (accessor *transformedAccessor) Int(obj interface{}) int {
	return accessor.fromAccessor.Int(obj)
}
