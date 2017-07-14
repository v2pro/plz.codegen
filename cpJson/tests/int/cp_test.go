package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cpJson"
	"github.com/google/gofuzz"
)

func Test_roundtrip(t *testing.T) {
	should := require.New(t)
	var src typeForTest
	fz := fuzz.New().MaxDepth(10).NilChance(0.3)
	fz.Fuzz(&src)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 512)
	should.Nil(plz.Copy(stream, src))
	var dst typeForTest
	iterator := jsoniter.ParseBytes(jsoniter.ConfigDefault, stream.Buffer())
	should.Nil(plz.Copy(&dst, iterator))
	should.Equal(src, dst)
}
