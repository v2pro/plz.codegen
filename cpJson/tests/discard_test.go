package tests

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cpJson/tests/model"
)

func Test_discard_extra_struct_field(t *testing.T) {
	should := require.New(t)

	dst := model.TypeE{}
	src := `{"Field1":"1","Field2":"2","Field3":"3"}`
	should.Nil(jsonCopy(&dst, src))
	should.Equal("1", dst.Field1)
	should.Equal("3", dst.Field3)
}

func Test_discard_extra_array_element(t *testing.T) {
	should := require.New(t)
	dst := [2]int{}
	src := `[1,2,3]`
	should.Nil(jsonCopy(&dst, src))
	should.Equal([2]int{1, 2}, dst)
}