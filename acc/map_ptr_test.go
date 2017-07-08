package acc

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"unsafe"
)

func Test_map_ptr_key_elem(t *testing.T) {
	should := require.New(t)
	v := map[int]*int{}
	should.Equal(lang.Int, objAcc(v).Key().Kind())
	should.Equal("*int", objAcc(v).Key().GoString())
	should.Equal(lang.Int, objAcc(v).Elem().Kind())
	should.Equal("**int", objAcc(v).Elem().GoString())
}

func Test_map_ptr_get_by_index(t *testing.T) {
	should := require.New(t)
	two := 2
	v := map[int]*int{1: &two}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	should.Equal(2, objAcc(v).Elem().Int(elem))
	elem = objAcc(v).MapIndex(objPtr(v), objPtr(100))
	should.Equal(unsafe.Pointer(nil), elem)
}

func Test_map_ptr_nil_get_by_index(t *testing.T) {
	should := require.New(t)
	v := map[int]*int{1: nil}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	should.Equal(unsafe.Pointer(nil), elem)
}

func Test_map_ptr_set_by_index(t *testing.T) {
	should := require.New(t)
	two := 2
	v := map[int]*int{1: &two}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	objAcc(v).Elem().SetInt(elem, 3)
	objAcc(v).SetMapIndex(objPtr(v), objPtr(1), elem)
	should.NotNil(v[1])
	should.Equal(3, *v[1])
}

func Test_map_ptr_iterate_map(t *testing.T) {
	should := require.New(t)
	two := 2
	v := map[int]*int{1: &two}
	keys := []int{}
	elems := []int{}
	objAcc(v).IterateMap(objPtr(v), func(key unsafe.Pointer, elem unsafe.Pointer) bool {
		keys = append(keys, objAcc(v).Key().Int(key))
		elems = append(elems, objAcc(v).Elem().Int(elem))
		return true
	})
	should.Equal([]int{1}, keys)
	should.Equal([]int{2}, elems)
}

func Test_map_ptr_nil_iterate_map(t *testing.T) {
	should := require.New(t)
	v := map[int]*int{1: nil}
	keys := []int{}
	elems := []int{}
	objAcc(v).IterateMap(objPtr(v), func(key unsafe.Pointer, elem unsafe.Pointer) bool {
		keys = append(keys, objAcc(v).Key().Int(key))
		if elem == nil {
			elems = append(elems, -1)
		} else {
			elems = append(elems, objAcc(v).Elem().Int(elem))
		}
		return true
	})
	should.Equal([]int{1}, keys)
	should.Equal([]int{-1}, elems)
}

func Test_map_ptr_fill_map(t *testing.T) {
	should := require.New(t)
	v := map[int]*int{}
	objAcc(v).FillMap(objPtr(v), func(filler lang.MapFiller) {
		key, elem := filler.Next()
		objAcc(v).Key().SetInt(key, 1)
		objAcc(v).Elem().SetInt(elem, 2)
		filler.Fill()
	})
	should.Equal(2, *v[1])
}
