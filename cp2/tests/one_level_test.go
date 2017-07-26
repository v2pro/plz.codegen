package tests

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	"reflect"
	"github.com/v2pro/wombat/cp2/tests/model"
)

func Test_ptr_int_2_int(t *testing.T) {
	runFuzzTest(t, generic.Int, reflect.PtrTo(generic.Int))
}

func Test_int_2_ptr_int(t *testing.T) {
	runFuzzTest(t, reflect.PtrTo(generic.Int), generic.Int)
}

func Test_array_int_2_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3]int{}), reflect.TypeOf([3]int{}))
}

func Test_array_int_2_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([]int{}), reflect.TypeOf([3]int{}))
}

func Test_slice_int_2_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3]int{}), reflect.TypeOf([]int{}))
}

func Test_slice_int_2_slice_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([]int{}), reflect.TypeOf([]int{}))
}

func Test_map_string_int_2_map_string_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string]int{}), reflect.TypeOf(map[string]int{}))
}

func Test_struct_2_struct(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(model.TypeA{}), reflect.TypeOf(model.TypeB{}))
}

func Test_struct_2_map_string_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf(map[string]int{}), reflect.TypeOf(model.TypeC{}))
}