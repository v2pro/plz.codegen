package cmpStructByField

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	f := Gen(reflect.TypeOf(TestObject{}), "Field")
	should.Equal(-1, f(
		TestObject{1}, TestObject{2}))
}

func by_reflect(obj1 interface{}, obj2 interface{}, fieldName string) int {
	field1 := reflect.ValueOf(obj1).FieldByName(fieldName).Int()
	field2 := reflect.ValueOf(obj2).FieldByName(fieldName).Int()
	if field1 < field2 {
		return -1
	} else if field1 == field2 {
		return 0
	} else {
		return 1
	}
}

func Benchmark_struct(b *testing.B) {
	type TestObject struct {
		Field int
	}
	f := Gen(reflect.TypeOf(TestObject{}), "Field")
	b.Run("plz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.ReportAllocs()
			f(TestObject{1}, TestObject{2})
		}
	})
	b.Run("reflect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.ReportAllocs()
			by_reflect(
				TestObject{1}, TestObject{2},
				"Field")
		}
	})
}
