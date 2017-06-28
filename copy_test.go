package wombat

import (
	"testing"
	"github.com/json-iterator/go/require"
)

func Test_copy(t *testing.T) {
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
