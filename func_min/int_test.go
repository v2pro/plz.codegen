package func_max

import (
	"testing"
	"github.com/v2pro/plz"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	should.Equal(1, plz.Min(1, 2, 3))
	should.Nil(plz.Min())
	reflect.ValueOf(1).Call()
}
