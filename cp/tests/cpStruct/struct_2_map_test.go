package cpStruct

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"testing"
)

func Test_to_map_new_entry(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	dst := map[string]int{}
	src := TestObject{100}
	f := cp.Gen(reflect.TypeOf(dst), reflect.TypeOf(src))
	should.Nil(f(dst, src))
	should.Equal(100, dst["Field"])
}

func Test_to_map_existing_entry(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field int
	}
	existing := int(0)
	dst := map[string]*int{"Field": &existing}
	src := TestObject{100}
	f := cp.Gen(reflect.TypeOf(dst), reflect.TypeOf(src))
	should.Nil(f(dst, src))
	should.Equal(100, existing)
}
