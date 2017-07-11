package gen

import (
	"testing"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_gen_simple_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(`
type fp_compare__TestObject struct {
	Field int
}`, generateStruct(reflect.TypeOf(TestObject{})))
}
