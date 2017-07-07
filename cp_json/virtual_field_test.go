package cp_json

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang/tagging"
	"github.com/v2pro/plz"
	"unsafe"
)

func Test_copy_from_virtual_field(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field1 string `json:"-"`
		Field2 string `json:"-"`
	}
	tagging.DefineStructTags(func(obj *TestObject) tagging.Tags {
		return tagging.D(
			tagging.S("codec", "json"),
			tagging.F(tagging.VirtualField("inner"), "json", tagging.V(
				"mapValue", func(ptr unsafe.Pointer) interface{} {
					obj := (*TestObject)(ptr)
					return struct{
						Field3 *string
						Field4 *string
					}{&obj.Field1, &obj.Field2}
				},
			)),
		)
	})

	obj := TestObject{"hello", "world"}
	output := []byte{}
	should.Nil(plz.Copy(&output, obj))
	should.Equal(`{"inner":{"Field3":"hello","Field4":"world"}}`, string(output))
}

func Test_copy_into_virtual_field(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field1 string `json:"-"`
		Field2 string `json:"-"`
	}
	tagging.DefineStructTags(func(obj *TestObject) tagging.Tags {
		return tagging.D(
			tagging.S("codec", "json"),
			tagging.F(tagging.VirtualField("inner"), "json", tagging.V(
				"mapValue", func(ptr unsafe.Pointer) interface{} {
					obj := (*TestObject)(ptr)
					return struct{
						Field3 *string
						Field4 *string
					}{&obj.Field1, &obj.Field2}
				},
			)),
		)
	})

	obj := TestObject{}
	should.Nil(plz.Copy(&obj, []byte(`{"inner":{"Field3":"hello","Field4":"world"}}`)))
	should.Equal("hello", obj.Field1)
	should.Equal("world", obj.Field2)
}