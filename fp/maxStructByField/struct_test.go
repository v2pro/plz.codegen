package max

import (
	"github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	f := genF(reflect.TypeOf(TestObject{}), "Field")
	should.Equal(TestObject{2}, f([]interface{}{
		TestObject{1}, TestObject{2},
	}))
}

func byReflect(objs []interface{}, fieldName string) interface{} {
	currentMax := reflect.ValueOf(objs[0])
	for i := 1; i < len(objs); i++ {
		elem := reflect.ValueOf(objs[i])
		if elem.FieldByName(fieldName).Int() > currentMax.FieldByName(fieldName).Int() {
			currentMax = elem
		}
	}
	return currentMax.Interface()
}

func Benchmark_struct(b *testing.B) {
	type TestObject struct {
		Field int
	}
	fuzzer := fuzz.New()
	datasets := make([][]interface{}, 32)
	for i := 0; i < len(datasets); i++ {
		dataset := make([]interface{}, 100)
		for j := 0; j < len(dataset); j++ {
			val := int(0)
			fuzzer.Fuzz(&val)
			dataset[j] = TestObject{val}
		}
		datasets[i] = dataset
	}
	f := genF(reflect.TypeOf(TestObject{}), "Field")
	b.Run("plz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.ReportAllocs()
			f(datasets[i%32])
		}
	})
	b.Run("reflect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.ReportAllocs()
			byReflect(datasets[i%32], "Field")
		}
	})
}
