package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_empty_interface_kind(t *testing.T) {
	should := require.New(t)
	var v []interface{}
	should.Equal(lang.Variant, objAcc(v).Elem().Kind())
}

func Test_empty_interface_gostring(t *testing.T) {
	should := require.New(t)
	var v []interface{}
	should.Equal("*interface {}", objAcc(v).Elem().GoString())
}

func Test_empty_interface_of_int_get_variant_elem(t *testing.T) {
	should := require.New(t)
	v := []interface{}{1}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Equal(1, realAcc.Int(realElem))
}

func Test_empty_interface_of_int_set_variant_elem(t *testing.T) {
	should := require.New(t)
	v := []interface{}{1}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Panics(func() {
		realAcc.SetInt(realElem, 2)
	})
}

func Test_empty_interface_of_ptr_int_get_variant_elem(t *testing.T) {
	should := require.New(t)
	v1 := int(1)
	v2 := &v1
	v := []interface{}{v2}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Equal(1, realAcc.Int(realElem))
}

func Test_empty_interface_of_ptr_int_set_variant_elem(t *testing.T) {
	should := require.New(t)
	v1 := int(1)
	v2 := &v1
	v := []interface{}{v2}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	realAcc.SetInt(realElem, 2)
	should.Equal(2, *(v[0].(*int)))
}

func Test_empty_interface_of_nil_get_variant_elem(t *testing.T) {
	should := require.New(t)
	v := []interface{}{nil}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Nil(realAcc)
	should.Equal(unsafe.Pointer(nil), realElem)
}
