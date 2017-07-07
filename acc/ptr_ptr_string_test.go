package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_ptr_ptr_string_kind(t *testing.T) {
	should := require.New(t)
	v1 := string("hello")
	v2 := &v1
	v := &v2
	should.Equal(lang.String, objAcc(v).Kind())
}

func Test_ptr_ptr_string_gostring(t *testing.T) {
	should := require.New(t)
	v1 := string("hello")
	v2 := &v1
	v := &v2
	should.Equal("**string", objAcc(v).GoString())
}

func Test_ptr_ptr_string_get_string(t *testing.T) {
	should := require.New(t)
	v1 := string("hello")
	v2 := &v1
	v := &v2
	should.Equal("hello", objAcc(v).String(objPtr(v)))
}

func Test_ptr_ptr_string_set_string(t *testing.T) {
	should := require.New(t)
	v1 := string("hello")
	v2 := &v1
	v := &v2
	objAcc(v).SetString(objPtr(v), "world")
	should.Equal("world", v1)
}

func Test_ptr_ptr_string_nil_set_string(t *testing.T) {
	should := require.New(t)
	var v1 *string
	v2 := &v1
	v := &v2
	objAcc(v).SetString(objPtr(v), "world")
	should.Equal("world", *v1)
}
