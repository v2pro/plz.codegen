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
	util.ObjectCopierProviders = append([]func(dstType, srcType reflect.Type) (util.ObjectCopier, error){
		provideFromRequestCopier,
	}, util.ObjectCopierProviders...)
}

func provideFromRequestCopier(dstType, srcType reflect.Type) (util.ObjectCopier, error) {
	isFromRequest := srcType == ptrHttpRequestType
	if !isFromRequest {
		return nil, nil
	}
	srcAcc := lang.AccessorOf(reflect.TypeOf(requestWrapper{}), "http")
	dstAcc := lang.AccessorOf(dstType, "http")
	copier, err := util.CopierOf(dstAcc, srcAcc)
	if err != nil {
		return nil, err
	}
	return &fromRequestCopier{copier}, nil
}

type fromRequestCopier struct {
	copier util.Copier
}

func (objCopier *fromRequestCopier) Copy(dst, src interface{}) error {
	req := src.(*http.Request)
	return objCopier.copier.Copy(lang.AddressOf(dst), lang.AddressOf(requestWrapper{req}))
}
