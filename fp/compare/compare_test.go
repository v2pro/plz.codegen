package compare

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/generic"
)

func Test_compare(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(F, "T", generic.Int).
	(func(int, int) int)
	should.Equal(-1, f(3, 4))
	should.Equal(0, f(3, 3))
	should.Equal(1, f(4, 3))
}