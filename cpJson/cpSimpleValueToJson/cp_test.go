package cpSimpleValueToJson

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	src := int(1)
	dst := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 512)
	should.Nil(plz.Copy(dst, src))
	should.Equal("1", string(dst.Buffer()))
}

func Test_int8(t *testing.T) {
	should := require.New(t)
	src := int8(1)
	dst := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 512)
	should.Nil(plz.Copy(dst, src))
	should.Equal("1", string(dst.Buffer()))
}