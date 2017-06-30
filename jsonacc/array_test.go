package jsonacc

import (
	"github.com/v2pro/plz"
	"testing"
	"github.com/json-iterator/go"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_array_decode(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
		`[1,2,3]`)
	accessor := plz.AccessorOf(reflect.TypeOf(iter))
	elems := []int{}
	accessor.IterateArray(iter, func(elem interface{}) bool {
		elems = append(elems, accessor.Elem().Int(elem))
		return true
	})
	should.Equal([]int{1, 2, 3}, elems)
}
