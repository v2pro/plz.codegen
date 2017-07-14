package cpStruct

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"github.com/v2pro/wombat/cp"
)

func Test_copy_struct_to_ptr_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := &TestObject{}
	src := TestObject{100}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(100, dst.Field)
}

func Test_copy_struct_to_ptr_nil_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := &TestObject{}
	src := TestObject{100}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	dst = nil
	should.Nil(f(&dst, src))
	should.Equal(100, dst.Field)
}
