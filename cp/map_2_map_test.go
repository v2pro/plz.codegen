package cp

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/wombat/cp/cpStatically"
	"reflect"
)

func Test_map_new_entry(t *testing.T) {
	should := require.New(t)
	dst := map[int]int{}
	src := map[int]int{1: 1, 2: 2}
	f := cpStatically.Gen(reflect.TypeOf(dst), reflect.TypeOf(src))
	should.Nil(f(dst, src))
	should.Equal(map[int]int{1: 1, 2: 2}, dst)
}

func Test_map_exiting_entry(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := map[int]*int{1: &existing}
	src := map[int]int{1: 1, 2: 2}
	f := cpStatically.Gen(reflect.TypeOf(dst), reflect.TypeOf(src))
	should.Nil(f(dst, src))
	should.Equal(1, existing)
	should.Equal(2, *dst[2])
}