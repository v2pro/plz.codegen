package jsonacc

import (
	"testing"
	"github.com/v2pro/plz"
	"reflect"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/require"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, "1")
	accessor := plz.AccessorOf(reflect.TypeOf(iter))
	should.Equal(1, accessor.Int(iter))
}
