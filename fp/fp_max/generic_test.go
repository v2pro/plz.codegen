package fp_max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
)

func Test_int8(t *testing.T) {
	should := require.New(t)
	should.Equal(int8(3), plz.Max(int8(1), int8(2), int8(3)))
}
