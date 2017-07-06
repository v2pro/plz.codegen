package cp

import (
	"testing"
	"github.com/v2pro/plz/util"
	"github.com/stretchr/testify/require"
)

func Test_copy_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	should.Nil(util.Copy(&dst, src))
	should.Equal(1, dst)
}

func Test_copy_int_to_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	should.NotNil(util.Copy(dst, src))
}

func Test_copy_string_to_ptr_string(t *testing.T) {
	should := require.New(t)
	dst := ""
	src := "world"
	should.Nil(util.Copy(&dst, src))
	should.Equal("world", dst)
}

func Test_copy_string_to_ptr_ptr_ptr_string(t *testing.T) {
	should := require.New(t)
	var v1 **string
	dst := &v1
	should.Nil(util.Copy(dst, "hello"))
	should.Equal("hello", ***dst)
}

func Test_copy_int_to_ptr_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	var v1 **int
	dst := &v1
	should.Nil(util.Copy(dst, 100))
	should.Equal(100, ***dst)
}
