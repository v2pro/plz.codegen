package cp_json

import (
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	_ "github.com/v2pro/wombat/cp"
	"reflect"
)

var iteratorType = reflect.TypeOf((*jsoniter.Iterator)(nil))
var streamType = reflect.TypeOf((*jsoniter.Stream)(nil))

func init() {
	lang.AccessorProviders = append([]func(typ reflect.Type) lang.Accessor{
		provideAccessor,
	}, lang.AccessorProviders...)
	util.ObjectCopierProviders = append([]func(dstType, srcType reflect.Type) (util.ObjectCopier, error){
		provideToBytesCopier,
		provideFromBytesCopier,
	}, util.ObjectCopierProviders...)
	util.CopierProviders = append([]func(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error){
		provideIteratorCopier,
		provideStreamCopier,
	}, util.CopierProviders...)
}

func provideAccessor(typ reflect.Type) lang.Accessor {
	if iteratorType == typ {
		return &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			lang.Variant,
		}
	}
	if streamType == typ {
		return &streamAccessor{
			lang.NoopAccessor{"streamAccessor"},
			lang.Variant,
		}
	}
	return nil
}

func provideStreamCopier(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error) {
	if srcAccessor.Kind() == lang.Variant {
		// use default impl
		return nil, nil
	}
	if dstAccessor.Kind() != lang.Variant {
		// already updated kind
		return nil, nil
	}
	dstStreamAccessor, isStream := dstAccessor.(*streamAccessor)
	if !isStream {
		return nil, nil
	}
	if srcAccessor.Kind() == lang.Struct {
		dstStreamAccessor = &streamAccessor{
			lang.NoopAccessor{"streamAccessor"},
			lang.Map,
		}
	} else {
		dstStreamAccessor = &streamAccessor{
			lang.NoopAccessor{"streamAccessor"},
			srcAccessor.Kind(),
		}
	}
	return util.CopierOf(dstStreamAccessor, srcAccessor)
}

func provideIteratorCopier(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error) {
	if dstAccessor.Kind() == lang.Variant {
		// use default impl
		return nil, nil
	}
	if srcAccessor.Kind() != lang.Variant {
		// already updated kind
		return nil, nil
	}
	srcIteratorAccessor, isIterator := srcAccessor.(*iteratorAccessor)
	if !isIterator {
		return nil, nil
	}
	if dstAccessor.Kind() == lang.Struct {
		srcIteratorAccessor = &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			lang.Map,
		}
	} else {
		srcIteratorAccessor = &iteratorAccessor{
			lang.NoopAccessor{"iteratorAccessor"},
			dstAccessor.Kind(),
		}
	}
	return util.CopierOf(dstAccessor, srcIteratorAccessor)
}
