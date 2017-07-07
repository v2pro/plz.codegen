package acc

import (
	"testing"
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang/tagging"
	"unsafe"
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
	should.Equal(0, acc.NumField())
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

func Test_field_virtual(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 string
	}
	tagging.DefineStructTags(func(obj *TestObject) tagging.Tags {
		return tagging.D(
			tagging.S(),
			tagging.F(tagging.VirtualField("field2"), "wombat", tagging.V(
				"mapValue", func(ptr unsafe.Pointer) interface{} {
					obj := (*TestObject)(ptr)
					return &obj.Field2
				},
			)),
		)
	})
	obj := TestObject{"hello", "world"}
	acc := lang.AccessorOf(reflect.TypeOf(obj), "wombat")
	should.Equal(3, acc.NumField())
	elem := acc.ArrayIndex(objPtr(obj), 2)
	should.Equal("world", acc.Field(2).Accessor().String(elem))
}
