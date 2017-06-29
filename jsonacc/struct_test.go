package jsonacc

import (
	"testing"
	"github.com/v2pro/plz"
	"reflect"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func Test_struct_decode(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	var v TestObject
	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
		`{"Field": 1}`)
	accessor := plz.AccessorOf(reflect.TypeOf(v), reflect.TypeOf(iter))
	should.Equal(reflect.Struct, accessor.Kind())
	should.Equal(reflect.String, accessor.Key().Kind())
	should.Equal(reflect.Interface, accessor.Elem().Kind())
	should.Equal(1, accessor.NumField())
	should.Equal("Field", accessor.Field(0).Name)
	elems := []int{}
	accessor.IterateMap(iter, func(key interface{}, elem interface{}) bool {
		elems = append(elems, accessor.Elem().Int(elem))
		return true
	})
}
