package fp_max

import (
	"testing"
	"github.com/v2pro/plz"
	"github.com/stretchr/testify/require"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 2, 3))
	should.Nil(plz.Max())
}
