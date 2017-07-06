package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
	"testing"
)

func Test_copy_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	should.Nil(util.Copy(&dst, src))
	should.Equal(1, dst)
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

func Test_copy_float64_to_ptr_float64(t *testing.T) {
	should := require.New(t)
	dst := float64(0)
	src := float64(1)
	should.Nil(util.Copy(&dst, src))
	should.Equal(float64(1), dst)
}
