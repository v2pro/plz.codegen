package wombat

import (
	"github.com/v2pro/plz"
	"reflect"
	_ "github.com/v2pro/wombat/jsonacc"
	_ "github.com/v2pro/plz/acc/native"
	"github.com/v2pro/wombat/cp"
)

func Copy(dst interface{}, src interface{}) error {
	dstAcc := plz.AccessorOf(reflect.TypeOf(dst))
	srcAcc := plz.AccessorOf(reflect.TypeOf(src))
	copier, err := cp.CopierOf(dstAcc, srcAcc)
	if err != nil {
		return err
	}
	return copier.Copy(dst, src)
}
