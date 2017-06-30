package jsonacc

import (
	"testing"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
	"reflect"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/acc"
)

func Test_map_decode(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
		`{"Field": 1}`)
	accessor := plz.AccessorOf(reflect.TypeOf(iter))
	should.Equal(acc.Interface, accessor.Kind())
	should.Equal(acc.String, accessor.Key().Kind())
	should.Equal(acc.Interface, accessor.Elem().Kind())
	keys := []string{}
	elems := []int{}
	accessor.IterateMap(iter, func(key interface{}, elem interface{}) bool {
		keys = append(keys, accessor.Key().String(key))
		elems = append(elems, accessor.Elem().Int(elem))
		return true
	})
	should.Equal([]string{"Field"}, keys)
	should.Equal([]int{1}, elems)
}