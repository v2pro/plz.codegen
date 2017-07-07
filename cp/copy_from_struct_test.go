package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
	"testing"
)

func Test_copy_struct_to_map(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 string
	}
	a := map[string]string{}
	should.Nil(util.Copy(a, TestObject{"1", "2"}))
	should.Equal(map[string]string{
		"Field1": "1",
		"Field2": "2",
	}, a)
}

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

func Test_copy_struct_to_ptr_struct_with_more_src(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
	}
	type B struct {
		Field  string
		Field2 string
	}
	var a A
	should.Nil(util.Copy(&a, B{"hello", "world"}))
	should.Equal("hello", a.Field)
}

func Test_copy_struct_to_ptr_struct_with_less_dst(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field  string
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

func Test_copy_struct_with_nil_ptr(t *testing.T) {
	should := require.New(t)
	type C struct {
		Field string
	}
	type B struct {
		Field1 *C
		Field2 *C
	}
	a := map[string]interface{}{}
	should.Nil(util.Copy(a, B{}))
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

func Test_copy_struct_with_map(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field map[string]string
	}
	type B struct {
		Field map[string]string
	}
	var a A
	should.Nil(util.Copy(&a, B{map[string]string{"hello": "world"}}))
	should.Equal(map[string]string{"hello": "world"}, a.Field)
}