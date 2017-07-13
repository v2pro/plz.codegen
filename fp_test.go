package wombat

import (
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/fp"
	"testing"
	"github.com/v2pro/wombat/gen"
	"github.com/stretchr/testify/require"
)

func init() {
	gen.CompilePlugin("/tmp/fp_test.so", func(){
		plz.Max(1, 3, 2)
	}, func(){
		plz.Max(
			User{1}, User{3}, User{2},
			"Score")
	})
	gen.LoadPlugin("/tmp/fp_test.so")
	gen.DisableDynamicCompilation()
}


type User struct {
	Score int
}
func Test_max_min(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	//should.Equal(1, plz.Min(1, 3, 2))
	should.Equal(User{3}, plz.Max(
		User{1}, User{3}, User{2},
		"Score"))
}
