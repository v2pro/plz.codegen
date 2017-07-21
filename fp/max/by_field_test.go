package max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
	"github.com/v2pro/wombat/generic"
	"github.com/json-iterator/go"
)

func Test_max_by_field(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}

	f := generic.Expand(ByField, "T", reflect.TypeOf([]TestObject{}), "F", "Field", "testMode", true).
	(func(interface{}) interface{})
	maxE := f([]TestObject{
		{1}, {3}, {2},
	})
	maxAsJson, err := jsoniter.MarshalToString(maxE)
	should.Nil(err)
	should.Equal(`{"Field":3}`, maxAsJson)
}
