package jsonacc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/json-iterator/go"
)

func init() {
	lang.AccessorProviders = append(lang.AccessorProviders, func(typ reflect.Type) lang.Accessor {
		if reflect.TypeOf((*jsoniter.Iterator)(nil)) == typ {
			return &iteratorAccessor{
				lang.NoopAccessor{"iteratorAccessor"},
				lang.Variant,
			}
		}
		if reflect.TypeOf((*jsoniter.Stream)(nil)) == typ {
			return &streamAccessor{
				lang.NoopAccessor{"streamAccessor"},
			}
		}
		return nil
	})
}
