package max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/testobj"
)

func Test_max_by_field(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(ByField, "T", reflect.TypeOf([]testobj.TestObject{}), "F", "Field").
	(func([]testobj.TestObject) testobj.TestObject)
	maxE := f([]testobj.TestObject{
		{1}, {3}, {2},
	})
	should.Equal(testobj.TestObject{3}, maxE)
}