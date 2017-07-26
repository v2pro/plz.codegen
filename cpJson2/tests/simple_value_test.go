package tests

import (
	"testing"
	"github.com/v2pro/wombat/generic"
)

func Test_int(t *testing.T) {
	runFuzzTest(t, generic.Int)
}

func Test_int8(t *testing.T) {
	runFuzzTest(t, generic.Int8)
}

func Test_int16(t *testing.T) {
	runFuzzTest(t, generic.Int16)
}

func Test_int32(t *testing.T) {
	runFuzzTest(t, generic.Int32)
}

func Test_int64(t *testing.T) {
	runFuzzTest(t, generic.Int64)
}

func Test_uint(t *testing.T) {
	runFuzzTest(t, generic.Uint)
}

func Test_uint8(t *testing.T) {
	runFuzzTest(t, generic.Uint8)
}

func Test_uint16(t *testing.T) {
	runFuzzTest(t, generic.Uint16)
}

func Test_uint32(t *testing.T) {
	runFuzzTest(t, generic.Uint32)
}

func Test_uint64(t *testing.T) {
	runFuzzTest(t, generic.Uint64)
}

func Test_float32(t *testing.T) {
	runFuzzTest(t, generic.Float32)
}

func Test_float64(t *testing.T) {
	runFuzzTest(t, generic.Float64)
}

func Test_string(t *testing.T) {
	runFuzzTest(t, generic.String)
}

func Test_bool(t *testing.T) {
	runFuzzTest(t, generic.Bool)
}

