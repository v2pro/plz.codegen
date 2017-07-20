package compare

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func Test_compare_ptr_int(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(comparePtr, "T", reflect.PtrTo(generic.Int)).
	(func(*int, *int) int)
	should.Equal(-1, f(ptrOf(3), ptrOf(4)))
	should.Equal(0, f(ptrOf(3), ptrOf(3)))
	should.Equal(1, f(ptrOf(4), ptrOf(3)))
}

func Test_compare_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(comparePtr, "T", reflect.PtrTo(reflect.PtrTo(generic.Int))).
	(func(**int, **int) int)
	should.Equal(-1, f(ptrPtrOf(3), ptrPtrOf(4)))
}

func ptrOf(obj int) *int {
	return &obj
}

func ptrPtrOf(obj int) **int {
	ptr := &obj
	return &ptr
}