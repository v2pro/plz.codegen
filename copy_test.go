package wombat

import (
	"testing"
	"github.com/json-iterator/go/require"
)

func Test_copy_struct_to_struct(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
	}
	type B struct {
		Field string
	}
	var a A
	should.Nil(Copy(&a, B{"hello"}))
	should.Equal("hello", a.Field)
}

func Test_copy_map_to_struct(t *testing.T) {
	should := require.New(t)
	type A struct {
		Field string
	}
	var a A
	b := map[string]string{
		"Field": "hello",
	}
	should.Nil(Copy(&a, b))
	should.Equal("hello", a.Field)
}

func Test_copy_struct_to_map(t *testing.T) {
	should := require.New(t)
	type B struct {
		Field string
	}
	b := B{"hello"}
	a := map[string]string{
	}
	should.Nil(Copy(a, b))
	should.Equal("hello", a["Field"])
}
