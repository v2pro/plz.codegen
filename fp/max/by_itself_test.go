package max

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_slice_int(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(ByItself, "T", reflect.TypeOf([]int{})).
	(func([]int) int)
	should.Equal(3, f([]int{1, 3, 2}))
}
