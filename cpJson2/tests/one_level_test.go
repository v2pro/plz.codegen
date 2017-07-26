package tests

import (
	"testing"
	"reflect"
)

func Test_array_int(t *testing.T) {
	runFuzzTest(t, reflect.TypeOf([3]int{}))
}