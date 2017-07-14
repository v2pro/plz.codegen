package cmpSimpleValue

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_ptr_int(t *testing.T) {
	should := require.New(t)
	f := genF(reflect.TypeOf(new(int)))
	one := int(1)
	zero := int(0)
	should.Equal(0, f(&one, &one))
	should.Equal(1, f(&one, &zero))
	should.Equal(-1, f(&zero, &one))
}

func Test_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	one := int(1)
	zero := int(0)
	ptrOne := &one
	ptrZero := &zero
	f := genF(reflect.TypeOf(&ptrZero))
	should.Equal(0, f(&ptrOne, &ptrOne))
	should.Equal(1, f(&ptrOne, &ptrZero))
	should.Equal(-1, f(&ptrZero, &ptrOne))
}
