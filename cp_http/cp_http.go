package cp_http

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"net/http"
	_ "github.com/v2pro/wombat/cp"
	"github.com/v2pro/plz/util"
)

var ptrHttpRequestType = reflect.TypeOf((*http.Request)(nil))

func init() {
	lang.AccessorProviders = append([]func(typ reflect.Type, tagName string) lang.Accessor{
		provideAccessor,
	}, lang.AccessorProviders...)
	util.ObjectCopierProviders = append([]func(dstType, srcType reflect.Type) (util.ObjectCopier, error){
		provideFromRequestCopier,
	}, util.ObjectCopierProviders...)
}

func provideAccessor(typ reflect.Type, tagName string) lang.Accessor {
	if ptrHttpRequestType == typ {
		return newRequestAccessor(tagName)
	}
	return nil
}

func provideFromRequestCopier(dstType, srcType reflect.Type) (util.ObjectCopier, error) {
	isFromRequest := srcType == ptrHttpRequestType
	if !isFromRequest {
		return nil, nil
	}
	srcAcc := lang.AccessorOf(reflect.TypeOf((*http.Request)(nil)), "http")
	dstAcc := lang.AccessorOf(dstType, "http")
	return util.DefaultObjectCopierOf(dstAcc, srcAcc)
}
