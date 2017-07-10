package acc

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
)

func Test_ptr_int8_kind(t *testing.T) {
	should := require.New(t)
	directV := int8(1)
	v := &directV
	should.Equal(lang.Int8, objAcc(v).Kind())
}

func Test_ptr8_int_gostring(t *testing.T) {
	should := require.New(t)
	directV := int8(1)
	v := &directV
	should.Equal("*int8", objAcc(v).GoString())
}

func Test_ptr_int8_get_int8(t *testing.T) {
	should := require.New(t)
	directV := int8(1)
	v := &directV
	should.Equal(int8(1), objAcc(v).Int8(objPtr(v)))
}

func Test_ptr_int8_set_int8(t *testing.T) {
	should := require.New(t)
	directV := int8(1)
	v := &directV
	objAcc(v).SetInt8(objPtr(v), 2)
	should.Equal(int8(2), objAcc(v).Int8(objPtr(v)))
}