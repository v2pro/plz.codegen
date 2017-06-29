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
	should.Equal(reflect.Map, accessor.Kind())
	elems := []int{}
	accessor.IterateMap(iter, func(key interface{}, elem interface{}) bool {
		elems = append(elems, accessor.Elem().Int(elem))
		return true
	})
}
