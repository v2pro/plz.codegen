package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_ptr_empty_interface_kind(t *testing.T) {
	should := require.New(t)
	var v *interface{}
	should.Equal(lang.Variant, objAcc(v).Kind())
}

func Test_ptr_empty_interface_gostring(t *testing.T) {
	should := require.New(t)
	var v *interface{}
	should.Equal("*interface {}", objAcc(v).GoString())
}

func Test_slice_empty_interface_of_int_get_variant_elem(t *testing.T) {
	should := require.New(t)
	v := &[]interface{}{1}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Equal(1, realAcc.Int(realElem))
}

func Test_slice_empty_interface_of_int_set_variant_elem(t *testing.T) {
	should := require.New(t)
	v := &[]interface{}{1}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().InitVariant(elem, objAcc(1))
	realAcc.SetInt(realElem, 2)
	should.Equal(2, (*v)[0].(int))
}

func Test_slice_empty_interface_of_nil_get_variant_elem(t *testing.T) {
	should := require.New(t)
	v := &[]interface{}{nil}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realAcc := objAcc(v).Elem().VariantElem(elem)
	should.Nil(realAcc)
	should.Equal(unsafe.Pointer(nil), realElem)
}

func Test_slice_empty_interface_of_nil_init_variant_elem(t *testing.T) {
	should := require.New(t)
	v := &[]interface{}{nil}
	elem := objAcc(v).ArrayIndex(objPtr(v), 0)
	realElem, realElemAcc := objAcc(v).Elem().InitVariant(elem, objAcc(1))
	realElemAcc.SetInt(realElem, 100)
	should.Equal(100, (*v)[0].(int))
}

func Test_ptr_empty_interface_of_int_get_variant_elem(t *testing.T) {
	should := require.New(t)
	var v1 interface{} = int(1)
	v := &v1
	realElem, realAcc := objAcc(v).VariantElem(objPtr(v))
	should.Equal(1, realAcc.Int(realElem))
}

func Test_ptr_empty_interface_of_int_set_variant_elem(t *testing.T) {
	should := require.New(t)
	var v1 interface{} = int(1)
	v := &v1
	realElem, realAcc := objAcc(v).InitVariant(objPtr(v), objAcc(1))
	realAcc.SetInt(realElem, 2)
	should.Equal(2, v1)
}

func Test_ptr_empty_interface_of_nil_get_variant_elem(t *testing.T) {
	should := require.New(t)
	var v1 interface{}
	v := &v1
	realElem, realAcc := objAcc(v).VariantElem(objPtr(v))
	should.Equal(unsafe.Pointer(nil), realElem)
	should.Nil(realAcc)
}

func Test_ptr_empty_interface_of_nil_set_variant_elem(t *testing.T) {
	should := require.New(t)
	var v1 interface{}
	v := &v1
	should.Nil(v1)
	realElem, realElemAcc := objAcc(v).InitVariant(objPtr(v), objAcc(123))
	realElemAcc.SetInt(realElem, 100)
	should.Equal(100, v1)
}
