package max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/google/gofuzz"
	"github.com/v2pro/wombat/gen"
	"reflect"
	"fmt"
)

func Test_src_struct(t *testing.T) {
	type TestObject struct {
		Field int
	}
	_, src := gen.Gen(F, "S", reflect.TypeOf(TestObject{}), "F", "Field", "T", reflect.TypeOf(int(0)))
	fmt.Println(src)
}

func Test_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(TestObject{2}, Call(
		[]interface{}{TestObject{1}, TestObject{2}},
		"Field"))
}

func by_reflect(objs []interface{}, fieldName string) interface{} {
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
	Call(datasets[0], "Field")
	b.Run("plz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.ReportAllocs()
			Call(datasets[i%32], "Field")
		}
	})
	b.Run("reflect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.ReportAllocs()
			by_reflect(datasets[i%32], "Field")
		}
	})
}
