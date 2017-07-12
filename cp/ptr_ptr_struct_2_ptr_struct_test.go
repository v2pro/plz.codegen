package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
	"testing"
)

func Test_copy_ptr_ptr_struct_2_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := TestObject{}
	src := TestObject{100}
	pSrc := &src
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&pSrc))
	should.Nil(f(&dst, &pSrc))
	should.Equal(100, dst.Field)
}

func Test_copy_ptr_nil_ptr_struct_2_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := TestObject{}
	src := TestObject{100}
	pSrc := &src
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&pSrc))
	pSrc = nil
	should.Nil(f(&dst, &pSrc))
	should.Equal(0, dst.Field)
}
