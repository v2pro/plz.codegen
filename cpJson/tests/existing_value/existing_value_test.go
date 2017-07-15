package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/json-iterator/go"
	_ "github.com/v2pro/wombat/cpJson"
)

func Test_array_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := [3]*int{&existing}
	src := `[1,2,3]`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func Test_slice_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := []*int{&existing}
	src := `[1,2,3]`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func Test_map_string_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := map[string]*int{"Field": &existing}
	src := `{"Field":1}`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func Test_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field *int
	}

	existing := int(0)
	dst := TestObject{Field:&existing}
	src := `{"Field":1}`
	should.Nil(jsonCopy(&dst, src))
	should.Equal(1, existing)
}

func jsonCopy(dst interface{}, srcJson string) error {
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, srcJson)
	return plz.Copy(dst, iter)
}
