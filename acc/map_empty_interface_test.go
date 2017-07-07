package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

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
