package cmpStructByField

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_struct_of_two_ptr(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 *int
		Field2 *int
	}
	f := genF(reflect.TypeOf(TestObject{}), "Field2")
	one := int(1)
	two := int(2)
	should.Equal(-1, f(
		TestObject{nil, &one}, TestObject{nil, &two}))
}

func Test_struct_of_one_ptr(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field *int
	}
	f := genF(reflect.TypeOf(TestObject{}), "Field")
	one := int(1)
	two := int(2)
	should.Equal(-1, f(
		TestObject{&one}, TestObject{&two}))
}
