package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cp"
)

func Test_ptr_ei_2_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	var src interface{} = int(1)
	should.Nil(plz.Copy(&dst, &src))
	should.Equal(1, dst)
}

func Test_ptr_ei_2_slice_int(t *testing.T) {
	should := require.New(t)
	dst := []int{}
	var src interface{} = []int{1, 2, 3}
	should.Nil(plz.Copy(&dst, &src))
	should.Equal([]int{1, 2, 3}, dst)
}

func Test_int_2_ptr_ei(t *testing.T) {
	should := require.New(t)
	var dst interface{}
	src := int(1)
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, dst)
}

func Test_slice_int_2_ptr_ei(t *testing.T) {
	should := require.New(t)
	var dst interface{}
	src := []int{1, 2, 3}
	should.Nil(plz.Copy(&dst, src))
	should.Equal([]int{1, 2, 3}, dst)
	src[0] = 2
	should.Equal([]int{1, 2, 3}, dst)
}

func Test_slice_ei_2_slice_int(t *testing.T) {
	should := require.New(t)
	src := []interface{}{1, 2, 3}
	dst := []int{}
	should.Nil(plz.Copy(&dst, src))
	should.Equal([]int{1, 2, 3}, dst)
}

func Test_slice_int_2_slice_ei(t *testing.T) {
	should := require.New(t)
	src := []int{1, 2, 3}
	dst := []interface{}{}
	should.Nil(plz.Copy(&dst, src))
	should.Equal([]interface{}{1, 2, 3}, dst)
}

func Test_slice_ei_2_slice_ei(t *testing.T) {
	should := require.New(t)
	src := []interface{}{1, "2", 3}
	dst := []interface{}{}
	should.Nil(plz.Copy(&dst, src))
	should.Equal([]interface{}{1, "2", 3}, dst)
}

func Test_map_string_ei_2_map_string_int(t *testing.T) {
	should := require.New(t)
	src := map[string]interface{}{"Field": 1}
	dst := map[string]int{}
	should.Nil(plz.Copy(dst, src))
	should.Equal(map[string]int{"Field": 1}, dst)
}

func Test_map_string_int_2_map_string_ei(t *testing.T) {
	should := require.New(t)
	src := map[string]int{"Field": 1}
	dst := map[string]interface{}{}
	should.Nil(plz.Copy(dst, src))
	should.Equal(map[string]interface{}{"Field": 1}, dst)
}

func Test_map_string_ei_2_map_string_ei(t *testing.T) {
	should := require.New(t)
	src := map[string]interface{}{"Field1": 1, "Field2": "2"}
	dst := map[string]interface{}{}
	should.Nil(plz.Copy(dst, src))
	should.Equal(map[string]interface{}{"Field1": 1, "Field2": "2"}, dst)
}
