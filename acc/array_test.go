package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_array_kind(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	should.Equal(lang.Array, objAcc(v).Kind())
}

func Test_array_gostring(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	should.Equal("[3]int", objAcc(v).GoString())
}

func Test_array_random_accessible(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	should.True(objAcc(v).RandomAccessible())
}

func Test_array_elem(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	should.Equal(lang.Int, objAcc(v).Elem().Kind())
}

func Test_array_get_by_array_index(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	should.Equal(2, objAcc(v).Elem().Int(elem))
}

func Test_array_get_by_array_index_out_of_bound(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), -1)
	})
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), 3)
	})
}

func Test_array_iterate_array(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Elem().Int(elem))
		return true
	})
	should.Equal([]int{1, 2, 3}, elems)
}

func Test_array_fill_array(t *testing.T) {
	should := require.New(t)
	v := [3]int{1, 2, 3}
	should.Panics(func() {
		objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		})
	})
}

func Test_array_of_ptr_int_with_nil(t *testing.T) {
	should := require.New(t)
	v := [3]*int{nil}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		should.Equal(unsafe.Pointer(nil), elem)
		return true
	})
}
