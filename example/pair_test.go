package example

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/example/model"
	"github.com/v2pro/wombat/container"
)

func Test_pair(t *testing.T) {
	should := require.New(t)
	intStringPairType := reflect.TypeOf(new(model.IntStringPair)).Elem()
	pair := generic.New(container.Pair, intStringPairType).(model.IntStringPair)
	should.Equal(0, pair.First())
	pair.SetFirst(1)
	should.Equal(1, pair.First())
}