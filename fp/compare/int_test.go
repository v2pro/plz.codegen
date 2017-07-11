package compare

import (
	"testing"
	"github.com/stretchr/testify/require"
	"reflect"
	"github.com/v2pro/wombat/gen"
)

func Test_src_int(t *testing.T) {
	should := require.New(t)
	_, src := gen.Gen(F, "T", reflect.TypeOf(int(0)))
	should.Equal(`
func Compare_int(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return typed_Compare_int(
		obj1.(int),
		obj2.(int))
}
func typed_Compare_int(
	obj1 int,
	obj2 int) int {
	// end of signature
	if (obj1 < obj2) {
		return -1
	} else if (obj1 == obj2) {
		return 0
	} else {
		return 1
	}
}`, src)
}

func Test_int(t *testing.T) {
	should := require.New(t)
	should.Equal(0, Call(1, 1))
	should.Equal(1, Call(1, 0))
	should.Equal(-1, Call(0, 1))
}

func Test_int8(t *testing.T) {
	should := require.New(t)
	should.Equal(0, Call(1, 1))
	should.Equal(0, Call(int8(1), int8(1)))
}

func Test_int16(t *testing.T) {
	should := require.New(t)
	should.Equal(0, Call(int16(1), int16(1)))
}
