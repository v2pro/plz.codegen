package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_copy_slice_to_slice(t *testing.T) {
	should := require.New(t)
	a := []int{}
	should.Nil(Copy(&a, []int{1, 2, 3}))
	should.Equal([]int{1, 2, 3}, a)
}
