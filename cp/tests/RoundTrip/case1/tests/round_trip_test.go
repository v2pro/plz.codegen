package cpVal

import (
	"github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_round_trip(t *testing.T) {
	should := require.New(t)
	dst := toType{}
	src := fromType{}
	fz := fuzz.New().MaxDepth(10).NilChance(0.3)
	fz.Fuzz(&src)
	f1 := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f1(&dst, src))
	f2 := cp.Gen(reflect.TypeOf(&src), reflect.TypeOf(dst))
	src2 := fromType{}
	should.Nil(f2(&src2, dst))
	should.Equal(src, src2)
}
