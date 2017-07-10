package acc

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
)

func Test_int8_kind(t *testing.T) {
	should := require.New(t)
	v := int8(1)
	should.Equal(lang.Int8, objAcc(v).Kind())
}

func Test_int8_gostring(t *testing.T) {
	should := require.New(t)
	v := int8(1)
	should.Equal("int8", objAcc(v).GoString())
}

func Test_int8_get_int8(t *testing.T) {
	should := require.New(t)
	v := int8(1)
	should.Equal(int8(1), objAcc(v).Int8(objPtr(v)))
}

func Test_int8_set_int8(t *testing.T) {
	should := require.New(t)
	v := int8(1)
	should.Panics(func() {
		objAcc(v).SetInt8(objPtr(v), 2)
	})
}
