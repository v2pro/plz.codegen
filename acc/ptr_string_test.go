package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_ptr_string_kind(t *testing.T) {
	should := require.New(t)
	directV := string("hello")
	v := &directV
	should.Equal(lang.String, objAcc(v).Kind())
}

func Test_ptr_string_gostring(t *testing.T) {
	should := require.New(t)
	directV := string("hello")
	v := &directV
	should.Equal("*string", objAcc(v).GoString())
}

func Test_ptr_string_get_string(t *testing.T) {
	should := require.New(t)
	directV := string("hello")
	v := &directV
	should.Equal("hello", objAcc(v).String(objPtr(v)))
}

func Test_ptr_string_set_string(t *testing.T) {
	should := require.New(t)
	directV := string("hello")
	v := &directV
	objAcc(v).SetString(objPtr(v), "world")
	should.Equal("world", directV)
}
