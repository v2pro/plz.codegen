package tests

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cp"
)

func Test_int_to_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	src := int(1)
	should.Nil(plz.Copy(&dst, src))
	should.Equal(int(1), dst)
}
