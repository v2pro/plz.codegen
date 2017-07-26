package tests

import (
	"testing"
	"github.com/v2pro/wombat/generic"
)

func Test_int(t *testing.T) {
	runFuzzTest(t, generic.Int)
}

func Test_string(t *testing.T) {
	runFuzzTest(t, generic.String)
}