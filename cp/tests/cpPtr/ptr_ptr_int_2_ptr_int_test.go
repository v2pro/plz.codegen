package cpPtr

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_copy_ptr_ptr_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	pSrc := &src
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&pSrc))
	should.Nil(f(&dst, &pSrc))
	should.Equal(1, dst)
}

func Test_copy_nil_ptr_ptr_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	pSrc := &src
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&pSrc))
	should.Nil(f(&dst, nil))
	should.Equal(0, dst)
}

func Test_copy_ptr_nil_ptr_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	pSrc := &src
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&pSrc))
	pSrc = nil
	should.Nil(f(&dst, &pSrc))
	should.Equal(0, dst)
}
