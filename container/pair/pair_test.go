package pair

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_pair(t *testing.T) {
	type IntStringPair interface {
		First() int
		SetFirst(val int)
		Second() string
		SetSecond(val string)
	}
	should := require.New(t)
	intStringPairType := reflect.TypeOf(new(IntStringPair)).Elem()
	pair := generic.New(Pair, intStringPairType).(IntStringPair)
	should.Equal(0, pair.First())
	pair.SetFirst(1)
	should.Equal(1, pair.First())
}