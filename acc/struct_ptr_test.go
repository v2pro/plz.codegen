package acc

import (
	"testing"
	"github.com/stretchr/testify/require"
	"unsafe"
)

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

func Test_struct_with_two_map(t *testing.T) {
	type TestObject struct {
		Field1 map[string]string
		Field2 map[string]string
	}
	should := require.New(t)
	v := TestObject{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		should.Equal(unsafe.Pointer(nil), elem)
		return true
	})
	should.Equal("*map[string]string", objAcc(v).Field(0).Accessor().GoString())
	v.Field2 = map[string]string{"hello": "world"}
	field2 := objAcc(v).ArrayIndex(objPtr(v), 1)
	mapAcc := objAcc(v).Field(1).Accessor()
	elem := mapAcc.MapIndex(field2, objPtr("hello"))
	should.Equal("world", mapAcc.Elem().String(elem))
}

func Test_struct_with_one_map(t *testing.T) {
	type TestObject struct {
		Field1 map[string]string
	}
	should := require.New(t)
	v := TestObject{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		should.Equal(unsafe.Pointer(nil), elem)
		return true
	})
	v.Field1 = map[string]string{"hello": "world"}
	field1 := objAcc(v).ArrayIndex(objPtr(v), 0)
	mapAcc := objAcc(v).Field(0).Accessor()
	elem := mapAcc.MapIndex(field1, objPtr("hello"))
	should.Equal("world", mapAcc.Elem().String(elem))
}