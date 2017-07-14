package case1

import (
	"testing"
	"github.com/google/gofuzz"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_round_trip(t *testing.T) {
	should := require.New(t)
	dst := fromType{}
	src := fromType{}
	fz := fuzz.New().MaxDepth(10).NilChance(0.3)
	fz.Fuzz(&dst)
	fz.Fuzz(&src)
	f := cp.Gen(reflect.TypeOf(&dst), reflect.TypeOf(src))
	should.Nil(f(&dst, src))
	should.Equal(dst, src)
}
