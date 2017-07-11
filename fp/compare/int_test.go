package compare

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	f := Gen(reflect.TypeOf(int(0)))
	should.Equal(0, f(1, 1))
	should.Equal(1, f(1, 0))
	should.Equal(-1, f(0, 1))
}

func Test_int8(t *testing.T) {
	should := require.New(t)
	f := Gen(reflect.TypeOf(int8(0)))
	should.Equal(0, f(int8(1), int8(1)))
}

func Test_int16(t *testing.T) {
	should := require.New(t)
	f := Gen(reflect.TypeOf(int16(0)))
	should.Equal(0, f(int16(1), int16(1)))
}
