package cpInterface

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func Test_empty_interface_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	src := int(1)
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf((*interface{})(nil)).Elem())
	should.Nil(f(&dst, src))
	should.Equal(int(1), dst)
}
