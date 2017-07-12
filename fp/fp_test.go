package fp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"testing"
)

func Test_max_of_nothing(t *testing.T) {
	should := require.New(t)
	should.Nil(plz.Max())
}

func Test_max_of_simple_value(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
}

func Test_max_of_struct_by_field(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(TestObject{2}, plz.Max(
		TestObject{1}, TestObject{2},
		"Field"))
}
