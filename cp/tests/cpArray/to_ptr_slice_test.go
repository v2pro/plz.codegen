package cpArray

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	_ "github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_to_ptr_slice(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := []int{}
	src := [3]int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([]int{1, 2, 3}, dst)
}
