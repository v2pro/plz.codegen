package cp_new

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
	"github.com/v2pro/wombat/cp_new/cp_statically"
)

func Test_copy_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	f := cp_statically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(1, dst)
}

func Test_copy_int8_to_ptr_int8(t *testing.T) {
	should := require.New(t)
	dst := int8(0)
	src := int8(1)
	f := cp_statically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(int8(1), dst)
}