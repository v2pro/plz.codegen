package fp_max

import (
	"testing"
	"github.com/v2pro/plz"
	"github.com/stretchr/testify/require"
)

func Test_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(TestObject{3}, plz.Max(
		TestObject{2}, TestObject{3}, TestObject{1},
		"Field"))
}
