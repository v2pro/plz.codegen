package cpPtr

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_copy_ptr_struct_2_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := TestObject{}
	src := TestObject{100}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(&src))
	should.Nil(f(&dst, &src))
	should.Equal(100, dst.Field)
}
