package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_slice_kind(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	should.Equal(lang.Array, objAcc(v).Kind())
}

func Test_slice_gostring(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	should.Equal("[]int", objAcc(v).GoString())
}

func Test_slice_random_accessible(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	should.True(objAcc(v).RandomAccessible())
}

func Test_slice_elem(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	should.Equal(lang.Int, objAcc(v).Elem().Kind())
}

func Test_slice_get_by_array_index(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	should.Equal(2, objAcc(v).Elem().Int(elem))
}

func Test_slice_get_by_array_index_out_of_bound(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), -1)
	})
	should.Panics(func() {
		objAcc(v).ArrayIndex(objPtr(v), 3)
	})
}

func Test_slice_iterate_array(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Elem().Int(elem))
		return true
	})
	should.Equal(v, elems)
}

func Test_slice_fill_array(t *testing.T) {
	should := require.New(t)
	v := []int{1, 2, 3}
	should.Panics(func() {
		objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		})
	})
}

func Test_slice_of_ptr_int(t *testing.T) {
	should := require.New(t)
	one := 1
	two := 2
	three := 3
	v := []*int{&one, &two, &three}
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Elem().Int(elem))
		return true
	})
	should.Equal([]int{1, 2, 3}, elems)
}

func Test_slice_of_ptr_int_with_nil(t *testing.T) {
	should := require.New(t)
	v := []*int{nil}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		should.Equal(unsafe.Pointer(nil), elem)
		return true
	})
}
