package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_ptr_int_kind(t *testing.T) {
	should := require.New(t)
	directV := int(1)
	v := &directV
	should.Equal(lang.Int, objAcc(v).Kind())
}

func Test_ptr_int_gostring(t *testing.T) {
	should := require.New(t)
	directV := int(1)
	v := &directV
	should.Equal("*int", objAcc(v).GoString())
}

func Test_ptr_int_get_int(t *testing.T) {
	should := require.New(t)
	directV := int(1)
	v := &directV
	should.Equal(1, objAcc(v).Int(objPtr(v)))
}

func Test_ptr_int_set_int(t *testing.T) {
	should := require.New(t)
	directV := int(1)
	v := &directV
	objAcc(v).SetInt(objPtr(v), 2)
	should.Equal(2, objAcc(v).Int(objPtr(v)))
}
