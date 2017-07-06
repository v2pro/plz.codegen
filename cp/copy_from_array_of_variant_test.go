package cp

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
)

func Test_copy_slice_of_variant_to_ptr_slice(t *testing.T) {
	should := require.New(t)
	a := []int{}
	should.Nil(util.Copy(&a, []interface{}{1, 2, 3}))
	should.Equal([]int{1, 2, 3}, a)
}

func Test_copy_slice_of_variant_to_ptr_slice_of_variant(t *testing.T) {
	should := require.New(t)
	a := []interface{}{}
	should.Nil(util.Copy(&a, []interface{}{1, 2, 3}))
	should.Equal([]interface{}{1, 2, 3}, a)
}
