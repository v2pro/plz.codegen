package cp_json

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/v2pro/plz/lang/tagging"
)

func Test_copy_into_bytes(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field string
	}
	tagging.Define(new(TestObject), "codec", "json")

	obj := TestObject{"hello"}
	output := []byte{}
	plz.Copy(&output, obj)
	should.Equal(`{"Field":"hello"}`, string(output))
}

func Test_copy_from_bytes(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field string
	}
	tagging.Define(new(TestObject), "codec", "json")

	obj := TestObject{}
	plz.Copy(&obj, []byte(`{"Field":"hello"}`))
	should.Equal("hello", obj.Field)
}
