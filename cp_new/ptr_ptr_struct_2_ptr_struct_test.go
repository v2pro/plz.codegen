package cp_new

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp_new/cp_statically"
	"reflect"
)

func Test_copy_ptr_ptr_struct_2_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := TestObject{}
	src := TestObject{100}
	pSrc := &src
	f := cp_statically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&pSrc))
	should.Nil(f(&dst, &pSrc))
	should.Equal(100, dst.Field)
}