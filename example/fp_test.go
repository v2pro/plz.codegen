package example

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/v2pro/wombat"
	_ "github.com/v2pro/wombat/fp"
	"testing"
	"github.com/v2pro/wombat/fp/maxSimpleValue"
	"github.com/v2pro/wombat/fp/maxStructByField"
	"reflect"
)

type User struct {
	Score int
}

func Test_max_min(t *testing.T) {
	wombat.Expand(maxSimpleValue.F,
		"T", wombat.Int)
	wombat.Expand(maxStructByField.F,
		"T", reflect.TypeOf(User{}),
		"F", "Score")
	wombat.CompilePlugin("/tmp/fp_test.so")
	wombat.LoadPlugin("/tmp/fp_test.so")
	wombat.DisableDynamicCompilation()

	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	should.Equal(User{3}, plz.Max(
		User{1}, User{3}, User{2},
		"Score"))
}
