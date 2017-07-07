package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_ptr_array_kind(t *testing.T) {
	should := require.New(t)
	directV := [3]int{}
	v := &directV
	should.Equal(lang.Array, objAcc(v).Kind())
}

func Test_ptr_array_gostring(t *testing.T) {
	should := require.New(t)
	directV := [3]int{}
	v := &directV
	should.Equal("*[3]int", objAcc(v).GoString())
}

func Test_ptr_array_random_accessible(t *testing.T) {
	should := require.New(t)
	directV := [3]int{}
	v := &directV
	should.True(objAcc(v).RandomAccessible())
}

func Test_ptr_array_get_by_array_index(t *testing.T) {
	should := require.New(t)
	directV := [3]int{1, 2, 3}
	v := &directV
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	should.Equal(2, objAcc(v).Elem().Int(elem))
}

func Test_ptr_array_set_by_array_index(t *testing.T) {
	should := require.New(t)
	directV := [3]int{1, 2, 3}
	v := &directV
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	objAcc(v).Elem().SetInt(elem, 4)
	should.Equal(4, directV[1])
}

func Test_ptr_array_iterate_array(t *testing.T) {
	should := require.New(t)
	directV := [3]int{1, 2, 3}
	v := &directV
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Elem().Int(elem))
		return true
	})
	should.Equal([]int{1, 2, 3}, elems)
}

func Test_ptr_array_fill_array(t *testing.T) {
	should := require.New(t)
	directV := [3]int{}
	v := &directV
	objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		index, elem := filler.Next()
		should.Equal(0, index)
		objAcc(v).Elem().SetInt(elem, 1)
		filler.Fill()
		index, elem = filler.Next()
		should.Equal(1, index)
		objAcc(v).Elem().SetInt(elem, 2)
		filler.Fill()
	})
	should.Equal([3]int{1, 2, 0}, directV)
}

func Test_ptr_array_fill_array_out_of_bound(t *testing.T) {
	should := require.New(t)
	directV := [0]int{}
	v := &directV
	objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		index, _ := filler.Next()
		should.Equal(-1, index)
	})
}
