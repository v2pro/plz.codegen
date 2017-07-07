package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_struct_kind(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Equal(lang.Struct, objAcc(v).Kind())
}

func Test_struct_gostring(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Equal("acc.TestObject", objAcc(v).GoString())
}

func Test_struct_num_field(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Equal(2, objAcc(v).NumField())
}

func Test_struct_field_name_tags(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Equal("Field1", objAcc(v).Field(0).Name())
	should.Equal(map[string]interface{}{
		"json": "field1",
	}, objAcc(v).Field(0).Tags())
	should.Equal("Field2", objAcc(v).Field(1).Name())
	should.Equal(map[string]interface{}{
		"json": "field2",
	}, objAcc(v).Field(1).Tags())
}

func Test_struct_field_accessor(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Equal(lang.Int, objAcc(v).Field(0).Accessor().Kind())
	should.Equal("*int", objAcc(v).Field(0).Accessor().GoString())
}

func Test_struct_random_accessible(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.True(objAcc(v).RandomAccessible())
}

func Test_struct_get_by_array_index(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v := TestObject{1, 2}
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	should.Equal(2, objAcc(v).Field(1).Accessor().Int(elem))
}

func Test_struct_get_by_array_index_out_of_bound(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), -1)
	})
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), 2)
	})
}

func Test_struct_iterate_array(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	should := require.New(t)
	v := TestObject{1, 2}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Field(index).Accessor().Int(elem))
		return true
	})
	should.Equal([]int{1, 2}, elems)
}

func Test_struct_with_string_field_iterate_array(t *testing.T) {
	type TestObject struct {
		Field1 string
		Field2 string
	}
	should := require.New(t)
	v := TestObject{"hello", "world"}
	elems := []string{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Field(index).Accessor().String(elem))
		return true
	})
	should.Equal([]string{"hello", "world"}, elems)
}

func Test_struct_fill_array(t *testing.T) {
	type TestObject struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}
	should := require.New(t)
	v := TestObject{1, 2}
	should.Panics(func() {
		objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		})
	})
}

func Test_struct_with_one_ptr(t *testing.T) {
	type TestObject struct {
		Field1 *int
	}
	should := require.New(t)
	one := 1
	v := TestObject{&one}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Field(index).Accessor().Int(elem))
		return true
	})
	should.Equal([]int{1}, elems)
}

func Test_struct_with_one_ptr_nil(t *testing.T) {
	type TestObject struct {
		Field1 *int
	}
	should := require.New(t)
	v := TestObject{nil}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		if elem == nil {
			elems = append(elems, -1)
		} else {
			elems = append(elems, objAcc(v).Field(index).Accessor().Int(elem))
		}
		return true
	})
	should.Equal([]int{-1}, elems)
}

func Test_struct_with_two_ptr_nil(t *testing.T) {
	type TestObject struct {
		Field1 *int
		Field2 *int
	}
	should := require.New(t)
	v := TestObject{nil, nil}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		if elem == nil {
			elems = append(elems, -1)
		} else {
			elems = append(elems, objAcc(v).Field(index).Accessor().Int(elem))
		}
		return true
	})
	should.Equal([]int{-1, -1}, elems)
}
