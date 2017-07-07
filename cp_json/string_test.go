package cp_json

import (
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	"reflect"
	"testing"
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

func Test_encode_ptr_string_nil(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
	should.Nil(util.Copy(stream, (*string)(nil)))
	should.Equal(`null`, string(stream.Buffer()))
}
