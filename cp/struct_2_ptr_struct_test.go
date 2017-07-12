package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
	"testing"
)

func Test_copy_same_type(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := TestObject{}
	src := TestObject{100}
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(100, dst.Field)
}

func Test_copy_different_type(t *testing.T) {
	should := require.New(t)

	type TestObject1 struct {
		Field int
	}
	type TestObject2 struct {
		Field int
	}
	dst := TestObject1{}
	src := TestObject2{100}
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(100, dst.Field)
}

func Test_copy_struct_of_ptr(t *testing.T) {
	should := require.New(t)

	type TestObject1 struct {
		Field int
	}
	type TestObject2 struct {
		Field *int
	}
	dst := TestObject1{}
	hundred := 100
	src := TestObject2{&hundred}
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(100, dst.Field)
}
