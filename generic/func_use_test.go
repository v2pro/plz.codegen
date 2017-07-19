package generic

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_expand_func_name(t *testing.T) {
	should := require.New(t)
	expanded := expandFuncName("cmp", []interface{}{"T", Int})
	should.Equal("cmp_T_int", expanded)
}
