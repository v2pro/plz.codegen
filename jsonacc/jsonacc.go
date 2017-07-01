package jsonacc

import (
	"github.com/v2pro/plz/acc"
	"reflect"
	"github.com/json-iterator/go"
)

func init() {
	acc.Providers = append(acc.Providers, func(typ reflect.Type) acc.Accessor {
		if reflect.TypeOf((*jsoniter.Iterator)(nil)) == typ {
			return &iteratorAccessor{}
		}
		if reflect.TypeOf((*jsoniter.Stream)(nil)) == typ {
			return &streamAccessor{}
		}
		return nil
	})
}
