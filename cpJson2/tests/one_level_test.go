package tests

import (
	"testing"
	"reflect"
)

func Test_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3]int{}))
}

func Test_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([]int{}))
}

func Test_map_string_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string]int{}))
}