package fp_max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
)

func Test_float64(t *testing.T) {
	should := require.New(t)
	should.Equal(2.1, plz.Max(1.0, 2.1, 2.0))
}
