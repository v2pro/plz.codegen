package fp_compare

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	should.Equal(0, Compare(1, 1))
	should.Equal(1, Compare(1, 0))
	should.Equal(-1, Compare(0, 1))
}

func Test_int8(t *testing.T) {
	should := require.New(t)
	should.Equal(0, Compare(1, 1))
	should.Equal(0, Compare(int8(1), int8(1)))
}

func Test_int16(t *testing.T) {
	should := require.New(t)
	should.Equal(0, Compare(int16(1), int16(1)))
}