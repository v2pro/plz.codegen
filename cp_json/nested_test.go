package cp_json

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang/tagging"
	"github.com/v2pro/plz"
)

func Test_encode_struct_of_map(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
		Mapped map[string]interface{}
	}
	tagging.Define(new(TestObject), "codec", "json")

	obj := TestObject{}
	output := []byte{}
	plz.Copy(&output, obj)
	should.Equal(`{"Field":0,"Mapped":null}`, string(output))
}
