package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
)

func Test_copy_nested_slice_to_slice(t *testing.T) {
	should := require.New(t)
	a := []interface{}{}
	should.Nil(Copy(&a, []interface{}{1, 2, []int{3, 4}}))
	should.Equal(1, a[0])
	should.Equal(2, a[1])
	should.Equal([]interface{}{3, 4}, a[2])
}

func Test_copy_nested_slice_to_json(t *testing.T) {
	should := require.New(t)
	a := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(Copy(a, []interface{}{1, 2, []int{3, 4}}))
	should.Equal("[1,2,[3,4]]", string(a.Buffer()))
}
