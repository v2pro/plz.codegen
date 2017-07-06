package jsoncp

import (
	"testing"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/v2pro/plz/util"
	"github.com/stretchr/testify/require"
)

func Test_decode_string_into_ptr_string(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, `"hello"`)
	accessor := lang.AccessorOf(reflect.TypeOf(iter))
	should.Equal(lang.Variant, accessor.Kind())
	val := ""
	should.Nil(util.Copy(&val, iter))
	should.Equal("hello", val)
}

func Test_decode_string_into_ptr_variant(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, `"hello"`)
	accessor := lang.AccessorOf(reflect.TypeOf(iter))
	should.Equal(lang.Variant, accessor.Kind())
	var val interface{}
	should.Nil(util.Copy(&val, iter))
	should.Equal("hello", val)
}

func Test_encode_string(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, "hello"))
	should.Equal(`"hello"`, string(stream.Buffer()))
}
