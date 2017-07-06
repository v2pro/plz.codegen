package jsoncp

import (
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
	"testing"
)

func Test_decode_map_into_map(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, `{"Field": 1}`)
	val := map[string]int{}
	should.Nil(util.Copy(val, iter))
	should.Equal(map[string]int{"Field": 1}, val)
}

func Test_decode_map_into_ptr_variant(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, `{"Field": 1}`)
	var val interface{}
	should.Nil(util.Copy(&val, iter))
	should.Equal(map[string]interface{}{"Field": float64(1)}, val)
}

func Test_decode_map_into_struct(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, `{"Field": 1}`)
	type TestObject struct {
		Field int
	}
	val := TestObject{}
	should.Nil(util.Copy(&val, iter))
	should.Equal(1, val.Field)
}

func Test_encode_map_of_string_to_int(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, map[string]int{"Field": 1}))
	should.Equal(`{"Field":1}`, string(stream.Buffer()))
}

func Test_encode_map_of_string_to_empty_interface(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, map[string]interface{}{"Field": 1}))
	should.Equal(`{"Field":1}`, string(stream.Buffer()))
}

func Test_encode_struct(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, &TestObject{1, 2}))
	should.Equal(`{"Field1":1,"Field2":2}`, string(stream.Buffer()))
}
