package cpVal

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_copy_val_to_val(t *testing.T) {
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
	pDst := &dst
	ppDst := &pDst
	f := cp.Gen(reflect.TypeOf(&ppDst), reflect.TypeOf(src))
	should.Nil(f(&ppDst, src))
	should.Equal(typeForTest(1), dst)
}

func Test_copy_int_to_ptr_nil_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	pDst := &dst
	ppDst := &pDst
	f := cp.Gen(reflect.TypeOf(&ppDst), reflect.TypeOf(src))
	ppDst = nil
	should.Nil(f(&ppDst, src))
	should.Equal(typeForTest(1), **ppDst)
}

func Test_copy_val_to_ptr_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	pDst := &dst
	f := cp.Gen(reflect.TypeOf(&pDst), reflect.TypeOf(src))
	should.Nil(f(&pDst, src))
	should.Equal(typeForTest(1), dst)
}

func Test_copy_val_to_nil_ptr_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	pDst := &dst
	f := cp.Gen(reflect.TypeOf(&pDst), reflect.TypeOf(src))
	// ignore nil dst
	should.Nil(f(nil, src))
}

func Test_copy_into_ptr_nil_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	pDst := &dst
	f := cp.Gen(reflect.TypeOf(&pDst), reflect.TypeOf(src))
	pDst = nil
	should.Nil(f(&pDst, src))
	should.Equal(typeForTest(1), *pDst)
}

func Test_copy_val_to_ptr_val(t *testing.T) {
	should := require.New(t)
	dst := typeForTest(0)
	src := typeForTest(1)
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(typeForTest(1), dst)
}
