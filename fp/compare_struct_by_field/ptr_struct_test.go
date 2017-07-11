package compare_struct_by_field

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_ptr_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	f := Gen(reflect.TypeOf(new(TestObject)), "Field")
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
	f := Gen(reflect.TypeOf(&one), "Field")
	should.Equal(-1, f(
		&one, &two))
}
