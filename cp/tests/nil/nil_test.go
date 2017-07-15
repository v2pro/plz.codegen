package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cp"
)

func Test_int_2_int(t *testing.T) {
	should := require.New(t)
	src := int(1)
	should.Nil(plz.Copy((*int)(nil), src))
}

func Test_ptr_int_2_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	should.Nil(plz.Copy(&dst, (*int)(nil)))
}

func Test_ptr_ptr_int_2_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	should.Nil(plz.Copy(&dst, (**int)(nil)))
	var src *int
	should.Nil(plz.Copy(&dst, &src))
}

func Test_ptr_int_2_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	pDst := &dst
	should.Nil(plz.Copy(&pDst, (*int)(nil)))
	should.Nil(pDst)
}

func Test_ptr_ptr_int_2_ptr_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	pDst := &dst
	should.Nil(plz.Copy(&pDst, (**int)(nil)))
	should.Nil(pDst)
	pDst = &dst
	var src *int
	should.Nil(plz.Copy(&pDst, &src))
	should.Nil(pDst)
}

func Test_slice_int_2_slice_int(t *testing.T) {
	should := require.New(t)
	dst := []int{1, 2, 3}
	var src []int
	should.Nil(plz.Copy(&dst, src))
	should.Nil(dst)
}

func Test_map_string_int_2_map_string_int(t *testing.T) {
	should := require.New(t)
	dst := map[string]int{"Field": 1}
	var src map[string]int
	should.Nil(plz.Copy(&dst, src))
	should.Nil(dst)
}
