package generic

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_expand_symbol_name(t *testing.T) {
	should := require.New(t)
	expanded := expandSymbolName("cmp", ArgMap{"T": Int})
	should.Equal("cmp_T_int", expanded)
}
