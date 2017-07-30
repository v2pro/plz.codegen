package example

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/fp"
	"testing"
	"github.com/v2pro/wombat/example/model"
	"github.com/v2pro/wombat/generic"
)

//go:generate go install github.com/v2pro/wombat/cmd/wombat-codegen
//go:generate $GOPATH/bin/wombat-codegen -pkg github.com/v2pro/wombat/example
func init() {
	generic.Declare(func() {
		plz.Max(int(0))
		plz.Max(float64(0))
		plz.Max(model.User{}, "Score")
	})
}

func Demo_max_min(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	should.Equal(float64(3), plz.Max(1.0, 3.0, 2.0))
	should.Equal(model.User{3}, plz.Max(
		model.User{1}, model.User{3}, model.User{2},
		"Score"))
}
