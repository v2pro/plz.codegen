package jsoncp

import (
	"testing"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
)

func Test_decode_array_into_ptr_slice(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
		`[1,2,3]`)
	elems := []int{}
	should.Nil(util.Copy(&elems, iter))
	should.Equal([]int{1, 2, 3}, elems)
}

func Test_decode_array_into_ptr_variant(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
		`[1,2,3]`)
	var elems interface{}
	should.Nil(util.Copy(&elems, iter))
	should.Equal([]interface{}{float64(1), float64(2), float64(3)}, elems)
}

func Test_encode_slice_of_int(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, []int{1, 2, 3}))
	should.Equal("[1,2,3]", string(stream.Buffer()))
}

func Test_encode_slice_of_variant(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, []interface{}{1, "2", 3}))
	should.Equal(`[1,"2",3]`, string(stream.Buffer()))
}

func Test_encode_slice_of_ptr_int(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	two := 2
	three := 3
	should.Nil(util.Copy(stream, []*int{nil, &two, &three}))
	should.Equal("[null,2,3]", string(stream.Buffer()))
}