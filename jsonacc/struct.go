package jsonacc

import (
	"reflect"
	"github.com/v2pro/plz/acc"
	"github.com/v2pro/plz"
	"github.com/json-iterator/go"
)

type structAccessor struct {
	mapAccessor
	typ reflect.Type
}

func (accessor *structAccessor) Kind() reflect.Kind {
	return reflect.Struct
}

func (accessor *structAccessor) GoString() string {
	return accessor.typ.Name()
}

func (accessor *structAccessor) NumField() int {
	return accessor.typ.NumField()
}

func (accessor *structAccessor) Field(index int) acc.StructField {
	field := accessor.typ.Field(index)
	return acc.StructField{
		Accessor: plz.AccessorOf(field.Type, reflect.TypeOf((*jsoniter.Iterator)(nil))),
		Name: field.Name,
		Tags: map[string]interface{}{},
	}
}
