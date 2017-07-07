package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_string_kind(t *testing.T) {
	should := require.New(t)
	v := string("hello")
	should.Equal(lang.String, objAcc(v).Kind())
}

func Test_string_gostring(t *testing.T) {
	should := require.New(t)
	v := string("hello")
	should.Equal("string", objAcc(v).GoString())
}

func Test_string_get_string(t *testing.T) {
	should := require.New(t)
	v := string("hello")
	should.Equal("hello", objAcc(v).String(objPtr(v)))
}

func Test_string_set_string(t *testing.T) {
	should := require.New(t)
	v := string("hello")
	should.Panics(func() {
		objAcc(v).SetString(objPtr(v), "world")
	})
}
