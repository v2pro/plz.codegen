package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
	"testing"
)

func Test_copy_ptr_ptr_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_src := &src
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&ptr_src))
	should.Nil(f(&dst, &ptr_src))
	should.Equal(1, dst)
}

func Test_copy_nil_ptr_ptr_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_src := &src
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&ptr_src))
	should.Nil(f(&dst, nil))
	should.Equal(0, dst)
}

func Test_copy_ptr_nil_ptr_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	ptr_src := &src
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&ptr_src))
	ptr_src = nil
	should.Nil(f(&dst, &ptr_src))
	should.Equal(0, dst)
}
