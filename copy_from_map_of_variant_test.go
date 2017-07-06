package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
)

func Test_copy_map_of_variant_to_map(t *testing.T) {
	should := require.New(t)
	a := map[string]string{}
	world := "world"
	should.Nil(util.Copy(&a, map[string]interface{}{"hello": &world}))
	should.Equal(map[string]string{"hello": "world"}, a)
}

func Test_copy_map_of_variant_to_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 string
	}
	a := TestObject{}
	should.Nil(util.Copy(&a, map[string]interface{}{"Field1": "hello", "Field2": "world"}))
	should.Equal("hello", a.Field1)
	should.Equal("world", a.Field2)
}

//func Test_copy_nested_struct_to_struct(t *testing.T) {
//	should := require.New(t)
//	type A1 struct {
//		Field2 string
//	}
//	type A struct {
//		Field1 A1
//	}
//	type B1 struct {
//		Field2 string
//	}
//	type B struct {
//		Field1 B1
//	}
//	var a A
//	should.Nil(Copy(&a, B{B1{"hello"}}))
//	should.Equal("hello", a.Field1.Field2)
//}
//
//func Test_copy_json_to_nested_struct(t *testing.T) {
//	should := require.New(t)
//	type A1 struct {
//		Field2 string
//	}
//	type A struct {
//		Field1 A1
//	}
//	b := jsoniter.ParseString(jsoniter.ConfigDefault, `{"Field1":{"Field2":"hello"}}`)
//	var a A
//	should.Nil(Copy(&a, b))
//	should.Equal("hello", a.Field1.Field2)
//}
//
//func Test_copy_nested_struct_to_map(t *testing.T) {
//	should := require.New(t)
//	type B1 struct {
//		Field2 string
//	}
//	type B struct {
//		Field1 B1
//	}
//	a := map[string]interface{}{}
//	should.Nil(Copy(a, B{B1{"hello"}}))
//	should.Equal("hello", a["Field1"].(B1).Field2)
//}
