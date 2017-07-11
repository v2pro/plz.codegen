package cp_new

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp_new/cp_statically"
	"reflect"
)

func Test_copy_int_to_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_dst := &dst
	f := cp_statically.Gen(reflect.TypeOf(&ptr_dst), reflect.TypeOf(src))
	should.Nil(f(&ptr_dst, src))
	should.Equal(1, dst)
}

func Test_copy_int_to_nil_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_dst := &dst
	f := cp_statically.Gen(reflect.TypeOf(&ptr_dst), reflect.TypeOf(src))
	// ignore nil dst
	should.Nil(f(nil, src))
}

func Test_copy_into_ptr_nil_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_dst := &dst
	f := cp_statically.Gen(reflect.TypeOf(&ptr_dst), reflect.TypeOf(src))
	ptr_dst = nil
	should.Nil(f(&ptr_dst, src))
	should.Equal(1, *ptr_dst)
}
