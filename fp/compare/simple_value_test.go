package compare

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	"github.com/stretchr/testify/require"
)

func Test_compare_int(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(compareSimpleValue, "T", generic.Int).
	(func(int, int) int)
	should.Equal(-1, f(3, 4))
	should.Equal(0, f(3, 3))
	should.Equal(1, f(4, 3))
}

func Test_compare_int_as_interface(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(compareSimpleValue, "T", generic.Int, "testMode", true).
	(func(interface{}, interface{}) int)
	should.Equal(-1, f(3, 4))
	should.Equal(0, f(3, 3))
	should.Equal(1, f(4, 3))
}
