package acc

import (
	"testing"
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/stretchr/testify/require"
)

func Test_ptr_float64_kind(t *testing.T) {
	should := require.New(t)
	directV := float64(1)
	v := &directV
	accessor := lang.AccessorOf(reflect.TypeOf(v))
	should.Equal(lang.Float64, accessor.Kind())
}

func Test_ptr_float64_gostring(t *testing.T) {
	should := require.New(t)
	directV := float64(1)
	v := &directV
	should.Equal("*float64", objAcc(v).GoString())
}

func Test_ptr_float64_get_float64(t *testing.T) {
	should := require.New(t)
	directV := float64(1)
	v := &directV
	should.Equal(float64(1), objAcc(v).Float64(objPtr(v)))
}

func Test_ptr_float64_set_float64(t *testing.T) {
	should := require.New(t)
	directV := float64(1)
	v := &directV
	objAcc(v).SetFloat64(objPtr(v), float64(2))
	should.Equal(float64(2), objAcc(v).Float64(objPtr(v)))
}

func Test_ptr_float64_nil_set_float64(t *testing.T) {
	should := require.New(t)
	var directV *float64
	v := &directV
	objAcc(v).SetFloat64(objPtr(v), float64(2))
	should.Equal(float64(2), objAcc(v).Float64(objPtr(v)))
}