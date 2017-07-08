package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_map_empty_interface_key_elem(t *testing.T) {
	should := require.New(t)
	v := map[int]interface{}{}
	should.Equal(lang.Int, objAcc(v).Key().Kind())
	should.Equal("*int", objAcc(v).Key().GoString())
	should.Equal(lang.Variant, objAcc(v).Elem().Kind())
	should.Equal("*interface {}", objAcc(v).Elem().GoString())
}

func Test_map_empty_interface_get_by_index(t *testing.T) {
	should := require.New(t)
	v := map[int]interface{}{1: 2}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Equal(2, realAcc.Int(realElem))
	elem = objAcc(v).MapIndex(objPtr(v), objPtr(100))
	should.Equal(unsafe.Pointer(nil), elem)
}

func Test_map_empty_interface_nil_get_by_index(t *testing.T) {
	should := require.New(t)
	v := map[int]interface{}{1: nil}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	should.Equal(unsafe.Pointer(nil), elem)
}

func Test_map_empty_interface_of_map_get_by_index(t *testing.T) {
	should := require.New(t)
	v := map[int]interface{}{1: map[int]int{2: 3}}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	should.NotNil(elem)
	realElem, realElemAcc := objAcc(v).Elem().VariantElem(elem)
	three := realElemAcc.MapIndex(realElem, objPtr(2))
	should.Equal(3, realElemAcc.Elem().Int(three))
}

func Test_map_empty_interface_set_by_index(t *testing.T) {
	should := require.New(t)
	v := map[int]interface{}{1: 2}
	elem := objAcc(v).MapIndex(objPtr(v), objPtr(1))
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	realAcc.SetInt(realElem, 3)
	objAcc(v).SetMapIndex(objPtr(v), objPtr(1), elem)
	should.Equal(map[int]interface{}{1: 3}, v)
}

func Test_map_empty_interface_iterate_map(t *testing.T) {
	should := require.New(t)
	v := map[string]interface{}{"hello": "world"}
	keys := []string{}
	elems := []string{}
	objAcc(v).IterateMap(objPtr(v), func(key unsafe.Pointer, elem unsafe.Pointer) bool {
		keys = append(keys, objAcc(v).Key().String(key))
		realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
		elems = append(elems, realAcc.String(realElem))
		return true
	})
	should.Equal([]string{"hello"}, keys)
	should.Equal([]string{"world"}, elems)
}

func Test_map_empty_interface_nil_iterate_map(t *testing.T) {
	should := require.New(t)
	v := map[string]interface{}{"hello": nil}
	keys := []string{}
	elems := []string{}
	objAcc(v).IterateMap(objPtr(v), func(key unsafe.Pointer, elem unsafe.Pointer) bool {
		keys = append(keys, objAcc(v).Key().String(key))
		if elem == nil {
			elems = append(elems, "nil")
		} else {
			realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
			elems = append(elems, realAcc.String(realElem))
		}
		return true
	})
	should.Equal([]string{"hello"}, keys)
	should.Equal([]string{"nil"}, elems)
}

func Test_map_empty_interface_fill_map(t *testing.T) {
	should := require.New(t)
	v := map[string]interface{}{}
	objAcc(v).FillMap(objPtr(v), func(filler lang.MapFiller) {
		key, elem := filler.Next()
		objAcc(v).Key().SetString(key, "hello")
		realElem, realAcc := objAcc(v).Elem().InitVariant(elem, objAcc(""))
		realAcc.SetString(realElem, "world")
		filler.Fill()
	})
	should.Equal(map[string]interface{}{"hello": "world"}, v)
}
