package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_copy_nested_struct_to_struct(t *testing.T) {
	should := require.New(t)
	type A1 struct {
		Field2 string
	}
	type A struct {
		Field1 A1
	}
	type B1 struct {
		Field2 string
	}
	type B struct {
		Field1 B1
	}
	var a A
	should.Nil(Copy(&a, B{B1{"hello"}}))
	should.Equal("hello", a.Field1.Field2)
}
