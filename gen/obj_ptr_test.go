package gen

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	f := objPtrGen(reflect.TypeOf(int(0)))
	should.Equal(100, *(*int)(f(100)))
}

func Test_ptr_int(t *testing.T) {
	should := require.New(t)
	hundred := int(100)
	f := objPtrGen(reflect.TypeOf(&hundred))
	should.Equal(100, *(*int)(f(&hundred)))
}

func Test_ptr_ptr_int(t *testing.T) {
	should := require.New(t)
	hundred := int(100)
	pHundred := &hundred
	f := objPtrGen(reflect.TypeOf(&pHundred))
	should.Equal(100, **(**int)(f(&pHundred)))
}

func Test_struct_of_one_ptr(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field *int
	}

	hundred := int(100)
	obj := TestObject{&hundred}
	f := objPtrGen(reflect.TypeOf(obj))
	should.Equal(100, *((*TestObject)(f(obj)).Field))
}
