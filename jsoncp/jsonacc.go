package jsoncp

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/json-iterator/go"
	_ "github.com/v2pro/wombat/cp"
	"github.com/v2pro/plz/util"
)

func init() {
	lang.AccessorProviders = append([]func(typ reflect.Type) lang.Accessor{
		provideAccessor,
	}, lang.AccessorProviders...)
	util.CopierProviders = append([]func(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error){
		provideCopier,
	}, util.CopierProviders...)
}

func provideAccessor(typ reflect.Type) lang.Accessor {
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
}

func provideCopier(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error) {
	if dstAccessor.Kind() == lang.Variant {
		// use default impl
		return nil, nil
	}
	srcIteratorAccessor, isIterator := srcAccessor.(*iteratorAccessor)
	if !isIterator {
		return nil, nil
	}
	if srcAccessor.Kind() != lang.Variant {
		// already updated kind
		return nil, nil
	}
	if dstAccessor.Kind() == lang.Struct {
		srcIteratorAccessor.kind = lang.Map
	} else {
		srcIteratorAccessor.kind = dstAccessor.Kind()
	}
	return util.CopierOf(dstAccessor, srcIteratorAccessor)
}
