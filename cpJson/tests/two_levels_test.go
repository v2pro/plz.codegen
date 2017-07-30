package tests

import (
	"testing"
	"reflect"
	"github.com/v2pro/wombat/cpJson/tests/model"
)

func Test_array_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3][3]int{}))
}

func Test_array_map_string_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3]map[string]int{}))
}

func Test_array_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3][]int{}))
}

func Test_array_struct(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3]model.TypeA{}))
}

func Test_slice_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([][3]int{}))
}

func Test_slice_map_string_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([]map[string]int{}))
}

func Test_slice_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([][]int{}))
}

func Test_slice_struct(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([]model.TypeA{}))
}

func Test_struct_and_other(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(model.TypeC{}))
}

func Test_map_string_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string][3]int{}))
}

func Test_map_string_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string][]int{}))
}

func Test_map_string_struct(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string]model.TypeA{}))
}

func Test_map_string_map_string_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string]map[string]int{}))
}