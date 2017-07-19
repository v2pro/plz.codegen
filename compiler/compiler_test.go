package compiler

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_dynamic_compile(t *testing.T) {
	should := require.New(t)
	plugin, err := DynamicCompile(`
package main
func Hello() string {
	return "Hello"
}
	`)
	should.Nil(err)
	symbol, err := plugin.Lookup("Hello")
	should.Nil(err)
	f := symbol.(func() string)
	should.Equal("Hello", f())
}