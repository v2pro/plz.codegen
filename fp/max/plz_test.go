package max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/v2pro/wombat/fp/testobj"
	"github.com/v2pro/wombat/generic"
)

func init() {
	generic.DynamicCompilationEnabled = true
}

func Test_max_of_nothing(t *testing.T) {
	should := require.New(t)
	should.Nil(plz.Max())
}

func Test_max_of_simple_value(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
}

func Test_max_of_struct_by_field(t *testing.T) {
	should := require.New(t)
	should.Equal(testobj.TestObject{2}, plz.Max(
		testobj.TestObject{1}, testobj.TestObject{2},
		"Field"))
}

