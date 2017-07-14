package cpInt

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"github.com/v2pro/wombat/cp"
)

func Test_copy_int_to_ptr_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_dst := &dst
	ptr_ptr_dst := &ptr_dst
	f := cp.Gen(reflect.TypeOf(&ptr_ptr_dst), reflect.TypeOf(src))
	should.Nil(f(&ptr_ptr_dst, src))
	should.Equal(1, dst)
}

func Test_copy_int_to_ptr_nil_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_dst := &dst
	ptr_ptr_dst := &ptr_dst
	f := cp.Gen(reflect.TypeOf(&ptr_ptr_dst), reflect.TypeOf(src))
	ptr_ptr_dst = nil
	should.Nil(f(&ptr_ptr_dst, src))
	should.Equal(1, **ptr_ptr_dst)
}
