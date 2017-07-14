package cpArray

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_array_more_element(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := [2]int{}
	src := [3]int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([2]int{1, 2}, dst)
}

func Test_array_less_element(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := [4]int{}
	src := [3]int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([4]int{1, 2, 3}, dst)
}
