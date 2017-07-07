package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"reflect"
	"testing"
)

func Test_int_kind(t *testing.T) {
	should := require.New(t)
	v := int(1)
	accessor := lang.AccessorOf(reflect.TypeOf(v))
	should.Equal(lang.Int, accessor.Kind())
}

func Test_int_gostring(t *testing.T) {
	should := require.New(t)
	v := int(1)
	should.Equal("int", objAcc(v).GoString())
}

func Test_int_get_int(t *testing.T) {
	should := require.New(t)
	v := int(1)
	should.Equal(1, objAcc(v).Int(objPtr(v)))
}

func Test_int_set_int(t *testing.T) {
	should := require.New(t)
	v := int(1)
	should.Panics(func() {
		objAcc(v).SetInt(objPtr(v), 2)
	})
}
