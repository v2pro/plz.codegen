package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_ptr_ptr_empty_interface_kind(t *testing.T) {
	should := require.New(t)
	var v **interface{}
	should.Equal(lang.Variant, objAcc(v).Kind())
}

func Test_ptr_ptr_empty_interface_gostring(t *testing.T) {
	should := require.New(t)
	var v **interface{}
	should.Equal("**interface {}", objAcc(v).GoString())
}

func Test_ptr_ptr_empty_interface_of_int_get_variant_elem(t *testing.T) {
	should := require.New(t)
	var v1 interface{} = int(1)
	v2 := &v1
	v := &v2
	realElem, realAcc := objAcc(v).VariantElem(objPtr(v))
	should.Equal(1, realAcc.Int(realElem))
}

func Test_ptr_ptr_empty_interface_of_int_set_variant_elem(t *testing.T) {
	should := require.New(t)
	var v1 interface{} = int(1)
	v2 := &v1
	v := &v2
	realElem, realAcc := objAcc(v).InitVariant(objPtr(v), objAcc(1))
	realAcc.SetInt(realElem, 2)
	should.Equal(2, v1)
}
