package tests

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/json-iterator/go"
	"github.com/v2pro/wombat/cpJson2/tests/model"
)

func Test_existing_array_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := [3]*int{&existing}
	src := `[1,2,3]`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func Test_existing_slice_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := []*int{&existing}
	src := `[1,2,3]`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func Test_existing_map_string_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := map[string]*int{"Field": &existing}
	src := `{"Field":1}`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func Test_existing_struct(t *testing.T) {
	should := require.New(t)

	existing := int(0)
	dst := model.TypeB{Field:&existing}
	src := `{"Field":1}`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func jsonCopy(dst interface{}, srcJson string) error {
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, srcJson)
	return plz.Copy(dst, iter)
}
