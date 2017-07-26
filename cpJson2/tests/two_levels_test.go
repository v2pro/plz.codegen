package tests

import (
	"testing"
	"reflect"
	"github.com/v2pro/wombat/cpJson2/tests/model"
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