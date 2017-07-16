package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	_ "github.com/v2pro/wombat/cpJson"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
)

func Test_discard_extra_struct_field(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field1 string
		Field3 string
	}

	dst := TestObject{}
	src := `{"Field1":"1","Field2":"2","Field3":"3"}`
	should.Nil(jsonCopy(&dst, src))
	should.Equal("1", dst.Field1)
	should.Equal("3", dst.Field3)
}

func Test_discard_extra_array_element(t *testing.T) {
	should := require.New(t)
	dst := [2]int{}
	src := `[1,2,3]`
	should.Nil(jsonCopy(&dst, src))
	should.Equal([2]int{1, 2}, dst)
}

func jsonCopy(dst interface{}, srcJson string) error {
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, srcJson)
	return plz.Copy(dst, iter)
}
