package fp_max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/google/gofuzz"
	"reflect"
)

func Test_int8(t *testing.T) {
	should := require.New(t)
	should.Equal(int8(3), plz.Max(int8(1), int8(2), int8(3)))
}

type simpleValueComparator interface {
	compare(obj1 reflect.Value, obj2 reflect.Value) int
}

type int8Comparator struct {
}

func (comparator *int8Comparator) compare(obj1 reflect.Value, obj2 reflect.Value) int {
	val1 := obj1.Int()
	val2 := obj2.Int()
	if val1 == val2 {
		return 0
	} else if val1 > val2 {
		return 1
	} else {
		return -1
	}
}

func max_comparator_reflect(comparator simpleValueComparator, collection []interface{}) interface{} {
	currentMax := reflect.ValueOf(collection[0])
	for _, elemObj := range collection[1:] {
		elem := reflect.ValueOf(elemObj)
		if comparator.compare(elem, currentMax) > 0 {
			currentMax = elem
		}
	}
	return currentMax.Interface()
}

func Benchmark_int8(b *testing.B) {
	fuzzer := fuzz.New()
	datasets := make([][]interface{}, 32)
	for i := 0; i < len(datasets); i++ {
		dataset := make([]interface{}, 100)
		for j := 0; j < len(dataset); j++ {
			val := int8(0)
			fuzzer.Fuzz(&val)
			dataset[j] = val
		}
		datasets[i] = dataset
	}
	var comparator simpleValueComparator = &int8Comparator{}
	b.Run("reflect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_comparator_reflect(comparator, datasets[i%32])
		}
	})
	b.Run("plz", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			plz.Max(datasets[i%32]...)
		}
	})
}