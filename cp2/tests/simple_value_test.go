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
