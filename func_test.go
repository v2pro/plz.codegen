package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
)

func Test_functional(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	should.Equal(1, plz.Min(1, 3, 2))

	type User struct {
		Score int
	}
	should.Equal(User{3}, plz.Max(User{1}, User{3}, User{2}, "Score"))
}