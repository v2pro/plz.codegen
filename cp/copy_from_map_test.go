package cp

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/util"
	"testing"
)

func Test_copy_nil_map_to_map(t *testing.T) {
	should := require.New(t)
	a := map[string]string{}
	var b *map[string]string
	should.Nil(util.Copy(a, &b))
	should.Equal(map[string]string{}, a)
}

func Test_copy_map_to_map(t *testing.T) {
	should := require.New(t)
	a := map[string]string{}
	should.Nil(util.Copy(a, map[string]string{"hello": "world"}))
	should.Equal(map[string]string{"hello": "world"}, a)
}

func Test_copy_map_to_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 string
	}
	a := TestObject{}
	should.Nil(util.Copy(&a, map[string]string{"Field1": "hello", "Field2": "world"}))
	should.Equal("hello", a.Field1)
	should.Equal("world", a.Field2)
}

func Test_copy_map_to_existing_map(t *testing.T) {
	should := require.New(t)
	existingValue := ""
	a := map[string]*string{
		"hello": &existingValue,
	}
	should.Nil(util.Copy(a, map[string]string{"hello": "world"}))
	should.Equal("world", existingValue)
}
