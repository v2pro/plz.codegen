package jsonacc

import (
	"testing"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_map_decode(t *testing.T) {
	should := require.New(t)
	v := map[string]int{}
	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
		`{"Field": 1}`)
	accessor := plz.AccessorOf(reflect.TypeOf(v), reflect.TypeOf(iter))
	should.Equal(reflect.Map, accessor.Kind())
	should.Equal(reflect.String, accessor.Key().Kind())
	should.Equal(reflect.Int, accessor.Elem().Kind())
	keys := []string{}
	elems := []int{}
	accessor.IterateMap(iter, func(key interface{}, elem interface{}) bool {
		keys = append(keys, accessor.Key().String(key))
		elems = append(elems, accessor.Elem().Int(elem))
		return true
	})
}