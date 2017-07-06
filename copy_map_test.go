package wombat
//
//import (
//	"testing"
//	"github.com/json-iterator/go/require"
//	"github.com/json-iterator/go"
//)
//
//func Test_copy_struct_to_struct(t *testing.T) {
//	should := require.New(t)
//	type A struct {
//		Field string
//	}
//	type B struct {
//		Field string
//	}
//	var a A
//	should.Nil(Copy(&a, B{"hello"}))
//	should.Equal("hello", a.Field)
//}
//
//func Test_copy_map_to_struct(t *testing.T) {
//	should := require.New(t)
//	type A struct {
//		Field string
//	}
//	var a A
//	b := map[string]string{
//		"Field": "hello",
//	}
//	should.Nil(Copy(&a, b))
//	should.Equal("hello", a.Field)
//}
//
//func Test_copy_struct_to_map(t *testing.T) {
//	should := require.New(t)
//	type B struct {
//		Field string
//	}
//	b := B{"hello"}
//	a := map[string]string{
//	}
//	should.Nil(Copy(a, b))
//	should.Equal("hello", a["Field"])
//}
//
//func Test_copy_json_to_map(t *testing.T) {
//	should := require.New(t)
//	b := jsoniter.ParseString(jsoniter.ConfigDefault, `{"Field":"hello"}`)
//	a := map[string]string{
//	}
//	should.Nil(Copy(a, b))
//	should.Equal("hello", a["Field"])
//}
//
//func Test_copy_json_to_struct(t *testing.T) {
//	should := require.New(t)
//	b := jsoniter.ParseString(jsoniter.ConfigDefault, `{"Field":"hello"}`)
//	type A struct {
//		Field string
//	}
//	var a A
//	should.Nil(Copy(&a, b))
//	should.Equal("hello", a.Field)
//}
//
//func Test_copy_struct_to_json(t *testing.T) {
//	should := require.New(t)
//	a := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
//	type B struct {
//		Field string
//	}
//	b := B{"hello"}
//	should.Nil(Copy(a, b))
//	should.Equal(`{"Field":"hello"}`, string(a.Buffer()))
//}
//
//func Test_copy_map_to_json(t *testing.T) {
//	should := require.New(t)
//	a := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
//	should.Nil(Copy(a, map[string]interface{}{
//		"hello": "world",
//	}))
//	should.Equal(`{"hello":"world"}`, string(a.Buffer()))
//}
