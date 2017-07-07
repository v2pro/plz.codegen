package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_ptr_ptr_struct_kind(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.Equal(lang.Struct, objAcc(v).Kind())
}

func Test_ptr_ptr_struct_gostring(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.Equal("**acc.TestObject", objAcc(v).GoString())
}

func Test_ptr_ptr_struct_num_field(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.Equal(2, objAcc(v).NumField())
}

func Test_ptr_ptr_struct_field_name_tags(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.Equal("Field1", objAcc(v).Field(0).Name())
	should.Equal(map[string]interface{}{
		"json": "field1",
	}, objAcc(v).Field(0).Tags())
	should.Equal("Field2", objAcc(v).Field(1).Name())
	should.Equal(map[string]interface{}{
		"json": "field2",
	}, objAcc(v).Field(1).Tags())
}

func Test_ptr_ptr_struct_field_accessor(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.Equal(lang.Int, objAcc(v).Field(0).Accessor().Kind())
	should.Equal("*int", objAcc(v).Field(0).Accessor().GoString())
}

func Test_ptr_ptr_struct_random_accessible(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.True(objAcc(v).RandomAccessible())
}

func Test_ptr_ptr_struct_get_by_array_index(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	should.Equal(2, objAcc(v).Field(1).Accessor().Int(elem))
}

func Test_ptr_ptr_struct_set_by_array_index(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	objAcc(v).Field(1).Accessor().SetInt(elem, 3)
	should.Equal(3, v1.Field2)
}

func Test_ptr_ptr_struct_get_by_array_index_out_of_bound(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), -1)
	})
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), 2)
	})
}

func Test_ptr_ptr_struct_iterate_array(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Field(index).Accessor().Int(elem))
		return true
	})
	should.Equal([]int{1, 2}, elems)
}

func Test_ptr_ptr_struct_fill_array(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v1 := &TestObject{1, 2}
	v := &v1
	objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		index, elem := filler.Next()
		should.Equal(0, index)
		objAcc(v).Field(index).Accessor().SetInt(elem, 2)
		filler.Fill()
		index, elem = filler.Next()
		should.Equal(1, index)
		objAcc(v).Field(index).Accessor().SetInt(elem, 3)
		filler.Fill()
	})
	should.Equal(2, v1.Field1)
	should.Equal(3, v1.Field2)
}
