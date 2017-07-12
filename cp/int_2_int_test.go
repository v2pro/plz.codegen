package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
	"testing"
)

func Test_copy_int_to_int(t *testing.T) {
	should := require.New(t)
	dst := 0
	src := 1
	should.Panics(func() {
		// int is not writable
		cpStatically.Gen(reflect.TypeOf(dst), reflect.TypeOf(src))
	})
}
