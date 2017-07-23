package tests

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func Test_ptr_int_2_int(t *testing.T) {
	runFuzzTest(t, generic.Int, reflect.PtrTo(generic.Int))
}