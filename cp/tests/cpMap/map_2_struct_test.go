package cpMap

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
)

func Test_to_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field1 int
		Field2 int
	}
	dst := TestObject{}
	src := map[string]int{
		"Field1": 100,
		"Field2": 200,
	}
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(100, dst.Field1)
	should.Equal(200, dst.Field2)
}