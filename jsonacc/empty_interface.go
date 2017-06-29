package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
)

type emptyInterfaceAccessor struct {
	acc.NoopAccessor
}

func (accessor *emptyInterfaceAccessor) Kind() reflect.Kind {
	return reflect.Interface
}

func (accessor *emptyInterfaceAccessor) GoString() string {
	return "interface{}"
}

func (accessor *emptyInterfaceAccessor) Int(obj interface{}) int {
	iter := obj.(*jsoniter.Iterator)
	return iter.ReadInt()
}

