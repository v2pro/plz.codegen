package jsoncp

import (
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/require"
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"reflect"
	"testing"
)

func Test_decode_number_into_ptr_int(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, "1")
	accessor := lang.AccessorOf(reflect.TypeOf(iter))
	should.Equal(lang.Variant, accessor.Kind())
	val := int(0)
	should.Nil(util.Copy(&val, iter))
	should.Equal(1, val)
}

func Test_decode_number_into_ptr_variant(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, "1")
	accessor := lang.AccessorOf(reflect.TypeOf(iter))
	should.Equal(lang.Variant, accessor.Kind())
	var val interface{}
	should.Nil(util.Copy(&val, iter))
	should.Equal(float64(1), val)
}

func Test_encode_int(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, 100))
	should.Equal("100", string(stream.Buffer()))
}

func Test_encode_variant_of_int(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	var val interface{} = 100
	should.Nil(util.Copy(stream, &val))
	should.Equal("100", string(stream.Buffer()))
}
