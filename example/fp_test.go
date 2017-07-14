package example

import (
	"github.com/v2pro/plz"
	"testing"
	"github.com/stretchr/testify/require"
	_ "github.com/v2pro/wombat/fp"
)

type User struct {
	Score int
}

func Test_max_min(t *testing.T) {
	//wombat.CompilePlugin("/tmp/fp_test.so", func() {
	//	plz.Max(1, 3, 2)
	//}, func() {
	//	plz.Max(
	//		User{1}, User{3}, User{2},
	//		"Score")
	//})
	//wombat.LoadPlugin("/tmp/fp_test.so")
	//wombat.DisableDynamicCompilation()

	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	should.Equal(User{3}, plz.Max(
		User{1}, User{3}, User{2},
		"Score"))
}
