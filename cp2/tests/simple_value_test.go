package tests

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	_ "github.com/v2pro/wombat/cp2"
)

func init() {
	generic.DynamicCompilationEnabled = true
}

func Test_int_2_int(t *testing.T) {
	runFuzzTest(t, generic.Int, generic.Int)
}

func Test_int8_2_int8(t *testing.T) {
	runFuzzTest(t, generic.Int8, generic.Int8)
}

func Test_int16_2_int16(t *testing.T) {
	runFuzzTest(t, generic.Int16, generic.Int16)
}

func Test_int32_2_int32(t *testing.T) {
	runFuzzTest(t, generic.Int32, generic.Int32)
}

func Test_int64_2_int64(t *testing.T) {
	runFuzzTest(t, generic.Int64, generic.Int64)
}

func Test_uint_2_uint(t *testing.T) {
	runFuzzTest(t, generic.Uint, generic.Uint)
}

func Test_uint8_2_uint8(t *testing.T) {
	runFuzzTest(t, generic.Uint8, generic.Uint8)
}

func Test_uint16_2_uint16(t *testing.T) {
	runFuzzTest(t, generic.Uint16, generic.Uint16)
}

func Test_uint32_2_uint32(t *testing.T) {
	runFuzzTest(t, generic.Uint32, generic.Uint32)
}

func Test_uint64_2_uint64(t *testing.T) {
	runFuzzTest(t, generic.Uint64, generic.Uint64)
}

func Test_bool_2_bool(t *testing.T) {
	runFuzzTest(t, generic.Bool, generic.Bool)
}

func Test_string_2_string(t *testing.T) {
	runFuzzTest(t, generic.String, generic.String)
}

func Test_float32_2_float32(t *testing.T) {
	runFuzzTest(t, generic.Float32, generic.Float32)
}

func Test_float64_2_float64(t *testing.T) {
	runFuzzTest(t, generic.Float64, generic.Float64)
}
