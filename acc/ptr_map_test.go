package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_ptr_map_kind(t *testing.T) {
	should := require.New(t)
	v := &map[int]int{1: 2}
	should.Equal(lang.Map, objAcc(v).Kind())
}

func Test_ptr_map_gostring(t *testing.T) {
	should := require.New(t)
	v := &map[int]int{1: 2}
	should.Equal("*map[int]int", objAcc(v).GoString())
}

func Test_ptr_map_key_elem(t *testing.T) {
	should := require.New(t)
	v := &map[int]int{1: 2}
	should.Equal(lang.Int, objAcc(v).Key().Kind())
	should.Equal("*int", objAcc(v).Key().GoString())
	should.Equal(lang.Int, objAcc(v).Elem().Kind())
	should.Equal("*int", objAcc(v).Elem().GoString())
}

func Test_ptr_map_iterate_map(t *testing.T) {
	should := require.New(t)
	v := &map[int]int{1: 2}
	keys := []int{}
	elems := []int{}
	objAcc(v).IterateMap(objPtr(v), func(key unsafe.Pointer, elem unsafe.Pointer) bool {
		keys = append(keys, objAcc(v).Key().Int(key))
		elems = append(elems, objAcc(v).Key().Int(elem))
		return true
	})
	should.Equal([]int{1}, keys)
	should.Equal([]int{2}, elems)
}

func Test_ptr_map_fill_map(t *testing.T) {
	should := require.New(t)
	v := &map[int]int{}
	objAcc(v).FillMap(objPtr(v), func(filler lang.MapFiller) {
		key, elem := filler.Next()
		objAcc(v).Key().SetInt(key, 1)
		objAcc(v).Elem().SetInt(elem, 2)
		filler.Fill()
	})
	should.Equal(map[int]int{1: 2}, *v)
}
