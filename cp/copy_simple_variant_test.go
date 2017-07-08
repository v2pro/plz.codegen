package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
	"testing"
)

func Test_copy_nil_nil_variant_to_ptr_string(t *testing.T) {
	should := require.New(t)
	a := ""
	var b **interface{}
	should.Nil(util.Copy(&a, &b))
	should.Equal("", a)
}

func Test_copy_nil_variant_to_ptr_string(t *testing.T) {
	should := require.New(t)
	a := ""
	var b interface{}
	should.Nil(util.Copy(&a, &b))
	should.Equal("", a)
}

func Test_copy_variant_to_ptr_string(t *testing.T) {
	should := require.New(t)
	a := ""
	var b interface{} = "hello"
	should.Nil(util.Copy(&a, &b))
	should.Equal("hello", a)
}

func Test_copy_ptr_variant_to_ptr_string(t *testing.T) {
	should := require.New(t)
	a := ""
	var b interface{} = "hello"
	c := &b
	should.Nil(util.Copy(&a, &c))
	should.Equal("hello", a)
}

func Test_copy_nil_to_variant(t *testing.T) {
	should := require.New(t)
	var a interface{}
	var b *string
	should.Nil(util.Copy(&a, b))
	should.Nil(b)
}

func Test_copy_string_to_variant(t *testing.T) {
	should := require.New(t)
	var a interface{}
	should.Nil(util.Copy(&a, "hello"))
	should.Equal("hello", a)
}

func Test_copy_ptr_variant_to_ptr_variant(t *testing.T) {
	should := require.New(t)
	var b interface{} = "hello"
	var a interface{}
	should.Nil(util.Copy(&a, &b))
	should.Equal("hello", a)
}

func Test_copy_ptr_variant_to_ptr_ptr_variant(t *testing.T) {
	should := require.New(t)
	var b interface{} = "hello"
	var a interface{}
	c := &a
	should.Nil(util.Copy(&c, &b))
	should.Equal("hello", a)
}
