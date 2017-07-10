package fp_max

import (
	"testing"
	"github.com/v2pro/plz"
	"github.com/stretchr/testify/require"
	"github.com/google/gofuzz"
	"reflect"
)

func Test_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(TestObject{3}, plz.Max(
		TestObject{2}, TestObject{3}, TestObject{1},
		"Field"))
}

type testObjectReflectComparator struct {
}

func (comparator *testObjectReflectComparator) compare(obj1 reflect.Value, obj2 reflect.Value) int {
	val1 := obj1.FieldByName("Field2").Int()
	val2 := obj2.FieldByName("Field2").Int()
	if val1 == val2 {
		return 0
	} else if val1 > val2 {
		return 1
	} else {
		return -1
	}
}

type TestObject struct {
	Field1 int
	Field2 int
}

func max_TestObject(collection ...interface{}) TestObject {
	currentMax := (collection[0]).(TestObject)
	for i := 1; i < len(collection); i++ {
		elem := collection[i].(TestObject)
		if compareTestObject(elem, currentMax) > 0 {
			currentMax = elem
		}
	}
	return currentMax
}

func compareTestObject(obj1 TestObject, obj2 TestObject) int {
	if obj1.Field2 < obj2.Field2 {
		return -1
	} else if obj1.Field2 < obj2.Field2 {
		return 0
	} else {
		return 1
	}
}

/*
reflect.value is 19x slower
2000000	       703 ns/op	       0 B/op	       0 allocs/op
100000	     19672 ns/op	    1568 B/op	     196 allocs/op
1000000	      1089 ns/op	       0 B/op	       0 allocs/op
 */
func Benchmark_struct(b *testing.B) {
	fuzzer := fuzz.New()
	datasets := make([][]interface{}, 32)
	comparableDatasets := make([][]interface{}, 32)
	typedDatasets := make([][]TestObject, 32)
	for i := 0; i < len(datasets); i++ {
		dataset := make([]interface{}, 100)
		comparableDataset := make([]interface{}, 100)
		typedDataset := make([]TestObject, 100)
		for j := 0; j < len(dataset); j++ {
			val := int(0)
			fuzzer.Fuzz(&val)
			dataset[j] = TestObject{0, val}
			comparableDataset[j] = TestObjectComparable{0, val}
			typedDataset[j] = TestObject{0, val}
		}
		dataset[99] = "Field2"
		datasets[i] = dataset
		comparableDatasets[i] = comparableDataset
		typedDatasets[i] = typedDataset
	}
	var comparator simpleValueComparator = &testObjectReflectComparator{}
	b.Run("typed", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_TestObject(datasets[i%32][:99]...)
		}
	})
	b.Run("comparable", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			plz.Max(comparableDatasets[i%32]...)
		}
	})
	b.Run("reflect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_comparator_reflect(comparator, datasets[i%32][:99])
		}
	})
	b.Run("plz", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			plz.Max(datasets[i%32]...)
		}
	})
}
