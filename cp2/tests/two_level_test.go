package tests

import (
	"testing"
	"reflect"
)

func Test_slice_slice_int_2_slice_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([][]int{}), reflect.TypeOf([][]int{}))
}