package jsonacc

import (
	"github.com/v2pro/plz"
	"testing"
	"github.com/json-iterator/go"
	"reflect"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/acc"
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

func Test_array_encode_one(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil , 1024)
	accessor := plz.AccessorOf(reflect.TypeOf(stream))
	accessor.FillArray(stream, func(filler acc.ArrayFiller) {
		elem := filler.Next()
		accessor.Elem().SetInt(elem, 1)
		filler.Fill()
	})
	should.Equal("[1]", string(stream.Buffer()))
}

func Test_array_encode_many(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil , 1024)
	accessor := plz.AccessorOf(reflect.TypeOf(stream))
	accessor.FillArray(stream, func(filler acc.ArrayFiller) {
		elem := filler.Next()
		accessor.Elem().SetInt(elem, 1)
		filler.Fill()
		elem = filler.Next()
		accessor.Elem().SetInt(elem, 2)
		filler.Fill()
	})
	should.Equal("[1,2]", string(stream.Buffer()))
}
