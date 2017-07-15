package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func Test_slice_have_less_elem(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := [4]int{}
	src := []int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([4]int{1, 2, 3}, dst)
}

func Test_slice_have_more_elem(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := [2]int{}
	src := []int{1, 2, 3}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal([2]int{1, 2}, dst)
}