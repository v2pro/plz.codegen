package acc

import (
	"testing"
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_field_rename(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string `wombat:"field"`
	}
	obj := TestObject{}
	acc := lang.AccessorOf(reflect.TypeOf(obj), "wombat")
	should.Equal(1, acc.NumField())
	should.Equal("field", acc.Field(0).Name())
}

func Test_field_skip(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string `wombat:"-"`
	}
	obj := TestObject{}
	acc := lang.AccessorOf(reflect.TypeOf(obj), "wombat")
	should.Equal(1, acc.NumField())
	should.Equal("", acc.Field(0).Name())
}

func Test_field_rename_to_dash(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string `wombat:"-,"`
	}
	obj := TestObject{}
	acc := lang.AccessorOf(reflect.TypeOf(obj), "wombat")
	should.Equal(1, acc.NumField())
	should.Equal("-", acc.Field(0).Name())
}

func Test_field_rename_to_original(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field string `wombat:","`
	}
	obj := TestObject{}
	acc := lang.AccessorOf(reflect.TypeOf(obj), "wombat")
	should.Equal(1, acc.NumField())
	should.Equal("Field", acc.Field(0).Name())
}
