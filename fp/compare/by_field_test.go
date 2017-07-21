package compare

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func Test_compare_by_field(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}

	testObjectType := reflect.TypeOf(TestObject{})
	f := generic.Expand(ByField, "T", testObjectType, "F", "Field", "testMode", true).
	(func(interface{}, interface{}) int)
	should.Equal(-1, f(TestObject{2}, TestObject{3}))
}