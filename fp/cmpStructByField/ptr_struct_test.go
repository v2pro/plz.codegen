package cmpStructByField

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_ptr_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	f := genF(reflect.TypeOf(new(TestObject)), "Field")
	should.Equal(-1, f(
		&TestObject{1}, &TestObject{2}))
}

func Test_ptr_ptr_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	one := &TestObject{1}
	two := &TestObject{2}
	f := genF(reflect.TypeOf(&one), "Field")
	should.Equal(-1, f(
		&one, &two))
}
