package tests

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
	"testing"
)

var fuzzTestFunc = generic.DefineFunc("RandomTest(t *testing.T)").
	Param("DT", "destination type").
	Param("ST", "source type").
	ImportPackage("testing").
	ImportPackage("github.com/stretchr/testify/require").
	ImportPackage("github.com/google/gofuzz").
	ImportPackage("github.com/v2pro/wombat/cp2/tests/helper").
	ImportPackage("github.com/v2pro/plz").
	Source(`
should := require.New(t)
fz := fuzz.New().MaxDepth(10).NilChance(0.3)
for i := 0; i < 100; i++ {
	var src {{.ST|name}}
	fz.Fuzz(&src)
	srcJson := helper.ToJson(src)
	var dst1 {{.DT|name}}
	should.Nil(plz.Copy(&dst1, src))
	var dst2 {{.DT|name}}
	helper.FromJson(&dst2, srcJson)
	dst1Json := helper.ToJson(dst1)
	dst2Json := helper.ToJson(dst2)
	should.Equal(dst2Json, dst1Json)
}`)

func runFuzzTest(t *testing.T, dstType reflect.Type, srcType reflect.Type) {
	f := generic.Expand(fuzzTestFunc, "DT", dstType, "ST", srcType).
	(func(*testing.T))
	f(t)
}