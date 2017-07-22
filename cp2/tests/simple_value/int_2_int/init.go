package test

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
)

func init() {
	var src SrcType
	generic.DeclareFunc(cp2.AnythingForPlz, "DT", reflect.TypeOf(new(DstType)), "ST", reflect.TypeOf(src))
}