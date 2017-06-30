package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
)

func Test_copy_slice_to_slice(t *testing.T) {
	should := require.New(t)
	a := []int{}
	should.Nil(Copy(&a, []int{1, 2, 3}))
	should.Equal([]int{1, 2, 3}, a)
}

func Test_copy_json_to_slice(t *testing.T) {
	should := require.New(t)
	b := jsoniter.ParseString(jsoniter.ConfigDefault, `[1,2,3]`)
	a := []int{}
	should.Nil(Copy(&a, b))
	should.Equal([]int{1, 2, 3}, a)
}
