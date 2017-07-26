package tests

import (
	"github.com/v2pro/wombat/generic"
	"testing"
	"reflect"
	_ "github.com/v2pro/wombat/cpJson2"
)

func init() {
	generic.DynamicCompilationEnabled = true
}

var fuzzTestFunc = generic.DefineFunc("RandomTest(t *testing.T)").
	Param("T", "type for test").
	ImportPackage("testing").
	ImportPackage("fmt").
	ImportPackage("reflect").
	ImportPackage("github.com/stretchr/testify/require").
	ImportPackage("github.com/google/gofuzz").
	ImportPackage("github.com/v2pro/plz").
	ImportPackage("github.com/json-iterator/go").
	Source(`
should := require.New(t)
fz := fuzz.New().MaxDepth(10).NilChance(0.3)
for i := 0; i < 100; i++ {
	var src {{.T|name}}
	fz.Fuzz(&src)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 512)
	should.Nil(plz.Copy(stream, src))
	var dst {{.T|name}}
	iterator := jsoniter.ParseBytes(jsoniter.ConfigDefault, stream.Buffer())
	should.Nil(plz.Copy(&dst, iterator))
	if !reflect.DeepEqual(src, dst) {
		fmt.Println(src)
		fmt.Println(string(stream.Buffer()))
		fmt.Println(dst)
		t.FailNow()
	}
}`)

func runFuzzTest(t *testing.T, typ reflect.Type) {
	f := generic.Expand(fuzzTestFunc, "T", typ).
	(func(*testing.T))
	f(t)
}
