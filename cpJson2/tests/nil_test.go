package tests

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
	"github.com/v2pro/wombat/cpJson2/tests/model"
)

func Test_nil_ptr_int(t *testing.T) {
	should := require.New(t)
	var src *int
	should.Equal("null", copyToJson(src))
	one := int(0)
	dst := &one
	should.Nil(copyFromJson(&dst, `null`))
	should.Nil(dst)
}

func Test_nil_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	var src *int
	pSrc := &src
	should.Equal(`null`, copyToJson(pSrc))
	pSrc = nil
	should.Equal(`null`, copyToJson(pSrc))
	var dst *int
	pDst := &dst
	should.Nil(copyFromJson(&pDst, `null`))
	should.Nil(pDst)
}

func Test_nil_slice_ptr_int(t *testing.T) {
	should := require.New(t)
	src := []*int{nil}
	should.Equal(`[null]`, copyToJson(src))
	one := int(1)
	dst := []*int{&one}
	should.Nil(copyFromJson(&dst, `[null]`))
	should.Nil(dst[0])
}

func Test_nil_slice_int(t *testing.T) {
	should := require.New(t)
	var src []int
	should.Equal(`null`, copyToJson(src))
	dst := []int{1}
	should.Nil(copyFromJson(&dst, `null`))
	should.Nil(dst)
}

func Test_nil_array_ptr_int(t *testing.T) {
	should := require.New(t)
	src := [3]*int{}
	should.Equal(`[null,null,null]`, copyToJson(src))
	one := int(1)
	dst := [3]*int{&one}
	should.Nil(copyFromJson(&dst, `[null]`))
	should.Nil(dst[0])
}

func Test_nil_map_string_ptr_int(t *testing.T) {
	should := require.New(t)
	src := map[string]*int{"Field": nil}
	should.Equal(`{"Field":null}`, copyToJson(src))
	one := int(1)
	dst := map[string]*int{"Field": &one}
	should.Nil(copyFromJson(&dst, `{"Field":null}`))
	should.Nil(dst["Field"])
}

func Test_nil_map_string_int(t *testing.T) {
	should := require.New(t)
	var src map[string]int
	should.Equal(`null`, copyToJson(src))
	dst := map[string]int{"Field": 1}
	should.Nil(copyFromJson(&dst, `null`))
	should.Nil(dst)
}

func Test_nil_struct(t *testing.T) {
	should := require.New(t)

	src := model.TypeB{}
	should.Equal(`{"Field":null}`, copyToJson(src))
	one := int(1)
	dst := model.TypeB{Field: &one}
	should.Nil(copyFromJson(&dst, `{"Field":null}`))
	should.Nil(dst.Field)
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
