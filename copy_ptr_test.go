package wombat
//
//import (
//	"testing"
//	"github.com/stretchr/testify/require"
//)
//
//func Test_copy_struct_of_ptr(t *testing.T) {
//	should := require.New(t)
//	type A struct {
//		Field *string
//	}
//	type B struct {
//		Field string
//	}
//	var a A
//	should.Nil(Copy(&a, B{"hello"}))
//	should.Equal("hello", a.Field)
//}