package cpSlice

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_slice_new_elem(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := []int{}
	src := []int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([]int{1, 2, 3}, dst)
}

func Test_slice_existing_elem(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	existing := int(0)
	dst := []*int{&existing}
	src := []int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(1, *dst[0])
	should.Equal(2, *dst[1])
	should.Equal(3, *dst[2])
}
