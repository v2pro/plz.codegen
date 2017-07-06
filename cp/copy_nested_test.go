package cp

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
)

func Test_copy_slice_of_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string
	}
	a := []TestObject{}
	should.Nil(util.Copy(&a, []TestObject{{"hello"}, {"world"}}))
	should.Equal(2, len(a))
	should.Equal("hello", a[0].Field)
	should.Equal("world", a[1].Field)
}

func Test_copy_slice_of_struct_into_slice_of_variant(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string
	}
	a := []interface{}{}
	should.Nil(util.Copy(&a, []TestObject{{"hello"}, {"world"}}))
	should.Equal(2, len(a))
	should.Equal(TestObject{"hello"}, a[0])
	should.Equal(TestObject{"world"}, a[1])
}

func Test_copy_slice_of_slice_into_slice_of_variant(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string
	}
	a := []interface{}{}
	should.Nil(util.Copy(&a, [][]int{{1, 2}, {3, 4}}))
	should.Equal(2, len(a))
	should.Equal([]int{1, 2}, a[0])
	should.Equal([]int{3, 4}, a[1])
}

func Test_copy_slice_of_variant_into_variant(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string
	}
	var a interface{}
	should.Nil(util.Copy(&a, []interface{}{TestObject{"hello"}, "world"}))
	should.Equal([]interface{}{TestObject{"hello"}, "world"}, a)
}
