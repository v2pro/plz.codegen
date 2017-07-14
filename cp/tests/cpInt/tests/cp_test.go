package cpVal

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_copy_int_to_int(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	should.Panics(func() {
		// int is not writable
		cp.Gen(reflect.TypeOf(dst), reflect.TypeOf(src))
	})
}

func Test_copy_val_to_ptr_ptr_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	ptr_dst := &dst
	ptr_ptr_dst := &ptr_dst
	f := cp.Gen(reflect.TypeOf(&ptr_ptr_dst), reflect.TypeOf(src))
	should.Nil(f(&ptr_ptr_dst, src))
	should.Equal(typeForTest(1), dst)
}

func Test_copy_int_to_ptr_nil_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	ptr_dst := &dst
	ptr_ptr_dst := &ptr_dst
	f := cp.Gen(reflect.TypeOf(&ptr_ptr_dst), reflect.TypeOf(src))
	ptr_ptr_dst = nil
	should.Nil(f(&ptr_ptr_dst, src))
	should.Equal(typeForTest(1), **ptr_ptr_dst)
}

func Test_copy_val_to_ptr_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	ptr_dst := &dst
	f := cp.Gen(reflect.TypeOf(&ptr_dst), reflect.TypeOf(src))
	should.Nil(f(&ptr_dst, src))
	should.Equal(typeForTest(1), dst)
}

func Test_copy_val_to_nil_ptr_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	ptr_dst := &dst
	f := cp.Gen(reflect.TypeOf(&ptr_dst), reflect.TypeOf(src))
	// ignore nil dst
	should.Nil(f(nil, src))
}

func Test_copy_into_ptr_nil_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	ptr_dst := &dst
	f := cp.Gen(reflect.TypeOf(&ptr_dst), reflect.TypeOf(src))
	ptr_dst = nil
	should.Nil(f(&ptr_dst, src))
	should.Equal(typeForTest(1), *ptr_dst)
}

func Test_copy_val_to_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(typeForTest(1), dst)
}
