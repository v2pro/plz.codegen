package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
	"testing"
)

func Test_copy_slice_to_ptr_slice(t *testing.T) {
	should := require.New(t)
	a := []int{}
	should.Nil(util.Copy(&a, []int{1, 2, 3}))
	should.Equal([]int{1, 2, 3}, a)
}

func Test_copy_array_to_ptr_slice(t *testing.T) {
	should := require.New(t)
	a := []int{}
	should.Nil(util.Copy(&a, [3]int{1, 2, 3}))
	should.Equal([]int{1, 2, 3}, a)
}

func Test_copy_slice_to_ptr_array(t *testing.T) {
	should := require.New(t)
	a := [1]int{}
	should.Nil(util.Copy(&a, []int{1, 2, 3}))
	should.Equal([]int{1}, a[:])
}

func Test_copy_array_to_array(t *testing.T) {
	should := require.New(t)
	a := [1]int{}
	should.Nil(util.Copy(&a, [3]int{1, 2, 3}))
	should.Equal([1]int{1}, a)
}

func Test_copy_slice_to_ptr_slice_of_variant(t *testing.T) {
	should := require.New(t)
	a := []interface{}{}
	should.Nil(util.Copy(&a, []int{1, 2, 3}))
	should.Equal([]interface{}{1, 2, 3}, a)
}

func Test_copy_slice_to_ptr_with_nil(t *testing.T) {
	should := require.New(t)
	a := []interface{}{}
	two := 2
	three := 3
	should.Nil(util.Copy(&a, []*int{nil, &two, &three}))
	should.Equal([]interface{}{nil, 2, 3}, a)
}
