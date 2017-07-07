package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_ptr_ptr_int_kind(t *testing.T) {
	should := require.New(t)
	v1 := int(1)
	v2 := &v1
	v := &v2
	should.Equal(lang.Int, objAcc(v).Kind())
}

func Test_ptr_ptr_int_gostring(t *testing.T) {
	should := require.New(t)
	v1 := int(1)
	v2 := &v1
	v := &v2
	should.Equal("**int", objAcc(v).GoString())
}

func Test_ptr_ptr_int_get_int(t *testing.T) {
	should := require.New(t)
	v1 := int(1)
	v2 := &v1
	v := &v2
	should.Equal(1, objAcc(v).Int(objPtr(v)))
}

func Test_ptr_ptr_int_set_int(t *testing.T) {
	should := require.New(t)
	v1 := int(1)
	v2 := &v1
	v := &v2
	objAcc(v).SetInt(objPtr(v), 2)
	should.Equal(2, v1)
}

func Test_ptr_ptr_int_nil_set_int(t *testing.T) {
	should := require.New(t)
	var v1 *int
	v2 := &v1
	v := &v2
	objAcc(v).SetInt(objPtr(v), 2)
	should.Equal(2, *v1)
}
