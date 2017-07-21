package example

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/fp"
	"testing"
	"github.com/v2pro/wombat/example/model"
)

func Test_max_min(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	should.Equal(model.User{3}, plz.Max(
		model.User{1}, model.User{3}, model.User{2},
		"Score"))
}
