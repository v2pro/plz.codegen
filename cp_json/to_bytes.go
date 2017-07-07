package cp_json

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"github.com/json-iterator/go"
	"unsafe"
	"reflect"
	"github.com/v2pro/plz/lang/tagging"
)

var ptrByteArrayType = reflect.TypeOf((*[]byte)(nil))

func provideToBytesCopier(dstType, srcType reflect.Type) (util.ObjectCopier, error) {
	isToBytes := dstType == ptrByteArrayType && tagging.Get(srcType).Tags["codec"] == "json"
	if !isToBytes {
		return nil, nil
	}
	dstAcc := lang.AccessorOf(reflect.TypeOf((*jsoniter.Stream)(nil)), "json")
	srcAcc := lang.AccessorOf(srcType, "json")
	copier, err := util.CopierOf(dstAcc, srcAcc)
	if err != nil {
		return nil, err
	}
	return &toBytesCopier{copier}, nil
}

type toBytesCopier struct {
	copier util.Copier
}

func (objCopier *toBytesCopier) Copy(dst, src interface{}) error {
	ptrBytes := dst.(*[]byte)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 512)
	err := objCopier.copier.Copy(unsafe.Pointer(stream), lang.AddressOf(src))
	if err != nil {
		return err
	}
	*ptrBytes = stream.Buffer()
	return nil
}
