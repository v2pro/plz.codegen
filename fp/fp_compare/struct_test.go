package fp_compare

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_struct(t *testing.T) {
	// TODO
	// 1. define struct
	// 2. include dependencies
	// 3. type unsafe cast
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(-1, CompareStructByField(TestObject{1}, TestObject{2}, "Field"))
}
