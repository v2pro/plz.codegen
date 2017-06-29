package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
)

type intAccessor struct {
	acc.NoopAccessor
}

func (accessor *intAccessor) Kind() reflect.Kind {
	return reflect.Int
}

func (accessor *intAccessor) GoString() string {
	return "int"
}

func (accessor *intAccessor) Int(obj interface{}) int {
	iter := obj.(*jsoniter.Iterator)
	return iter.ReadInt()
}
