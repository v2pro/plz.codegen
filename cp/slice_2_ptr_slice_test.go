package cp

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
)

func Test_slice_to_ptr_slice(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := []int{}
	src := []int{1, 2, 3}
	f := cpStatically.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([]int{1, 2, 3}, dst)
}
