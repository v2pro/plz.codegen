package existing_value

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cp"
)

func Test_int_2_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := int(1)
	pDst := &dst
	src := int(2)
	should.Nil(plz.Copy(&pDst, src))
	should.Equal(2, dst)
}

func Test_slice_int_2_slice_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := []*int{&existing}
	src := []int{1, 2, 3}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, existing)
	should.Equal(2, *dst[1])
	should.Equal(3, *dst[2])
}

func Test_array_int_2_array_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := [3]*int{&existing}
	src := [3]int{1, 2, 3}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, existing)
	should.Equal(2, *dst[1])
	should.Equal(3, *dst[2])
}

func Test_struct_int_2_struct_ptr_int(t *testing.T) {
	should := require.New(t)
	type srcType struct {
		Field int
	}
	type dstType struct {
		Field *int
	}
	existing := int(0)
	dst := dstType{Field: &existing}
	src := srcType{Field: 1}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, existing)
}

func Test_map_string_int_2_map_string_ptr_int(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	dst := map[string]*int{"Field": &existing}
	src := map[string]int{"Field": 1}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, existing)
}

func Test_int_2_ptr_ei(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	var dst interface{} = &existing
	src := int(1)
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, existing)
}

func Test_slice_int_2_ptr_ei(t *testing.T) {
	should := require.New(t)
	existing := int(0)
	var dst interface{} = &[]*int{&existing}
	src := []int{1, 2, 3}
	should.Nil(plz.Copy(&dst, src))
	should.Equal(1, existing)
}