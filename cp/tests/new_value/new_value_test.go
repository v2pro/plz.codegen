package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cp"
)

func Test_int_2_ptr_int(t *testing.T) {
	should := require.New(t)
	var dst *int
	src := int(1)
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, *dst)
}

func Test_array_int_2_ptr_array_int(t *testing.T) {
	should := require.New(t)
	var dst *[3]int
	src := [3]int{1, 2, 3}
	should.Nil(plz.Copy(&dst, src))
	should.Equal([3]int{1, 2, 3}, *dst)
}

func Test_struct_2_ptr_struct(t *testing.T) {
	should := require.New(t)

	type TestObject struct{
		Field int
	}

	var dst *TestObject
	src := TestObject{Field:1}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, dst.Field)
}

func Test_slice_int_2_slice_int(t *testing.T) {
	should := require.New(t)
	var dst []int
	src := []int{1, 2, 3}
	should.Nil(plz.Copy(&dst, src))
	should.Equal([]int{1, 2, 3}, dst)
}

func Test_map_string_int_2_map_string_int(t *testing.T) {
	should := require.New(t)
	var dst map[string]int
	src := map[string]int{"Field": 1}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(map[string]int{"Field": 1}, dst)
}
