package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
)

func Test_copy_struct_to_ptr_struct_with_exact_fields(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
	}
	type B struct {
		Field string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello"}))
	should.Equal("hello", a.Field)
}

func Test_copy_struct_to_struct(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
	}
	type B struct {
		Field string
	}
	var a A
	should.NotNil(util.Copy(a, B{"hello"}))
}

func Test_copy_struct_to_ptr_struct_with_more_src(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
	}
	type B struct {
		Field string
		Field2 string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello", "world"}))
	should.Equal("hello", a.Field)
}

func Test_copy_struct_to_ptr_struct_with_less_dst(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
		Field2 string
	}
	type B struct {
		Field string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello"}))
	should.Equal("hello", a.Field)
}

func Test_copy_struct_to_ptr_struct_with_no_match(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field1 string
	}
	type B struct {
		Field2 string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello"}))
	should.Equal("", a.Field1)
}

func Test_copy_struct_with_ptr(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field *string
	}
	type B struct {
		Field string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello"}))
	should.Equal("hello", *a.Field)
}

func Test_copy_struct_with_ptr_ptr(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field **string
	}
	type B struct {
		Field string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello"}))
	should.Equal("hello", **a.Field)
}
