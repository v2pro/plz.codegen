package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cpJson"
)

func Test_ptr_int(t *testing.T) {
	should := require.New(t)
	var src *int
	should.Equal("null", copyToJson(src))
	one := int(0)
	dst := &one
	should.Nil(copyFromJson(&dst, `null`))
	should.Nil(dst)
}

func Test_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	var src *int
	pSrc := &src
	should.Equal("null", copyToJson(pSrc))
	pSrc = nil
	should.Equal("null", copyToJson(pSrc))
	var dst *int
	pDst := &dst
	should.Nil(copyFromJson(&pDst, `null`))
	should.Nil(pDst)
}

func copyToJson(src interface{}) string {
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 512)
	err := plz.Copy(stream, src)
	if err != nil {
		panic(err.Error())
	}
	return string(stream.Buffer())
}

func copyFromJson(dst interface{}, srcJson string) error {
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, srcJson)
	return plz.Copy(dst, iter)
}
