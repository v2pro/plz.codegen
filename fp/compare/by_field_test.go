package compare

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/generic"
	"reflect"
	"github.com/v2pro/wombat/fp/testobj"
)

func Test_compare_by_field(t *testing.T) {
	should := require.New(t)
	testObjectType := reflect.TypeOf(testobj.TestObject{})
	f := generic.Expand(ByField, "T", testObjectType, "F", "Field").
	(func(testobj.TestObject, testobj.TestObject) int)
	should.Equal(-1, f(testobj.TestObject{2}, testobj.TestObject{3}))
}