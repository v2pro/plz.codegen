package fp_max

import (
	"testing"
	"github.com/v2pro/plz"
	"github.com/stretchr/testify/require"
	"github.com/google/gofuzz"
	"reflect"
	"math"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 2, 3))
	should.Nil(plz.Max())
}

const maxUint = ^uint(0)
const minInt = -int(maxUint >> 1)

func max_int_typed(collection ...int) int {
	currentMax := minInt
	for _, elem := range collection {
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}

func max_int_manual(collection ...interface{}) interface{} {
	currentMax := minInt
	for _, elem := range collection {
		elemVal := elem.(int)
		if elemVal > currentMax {
			currentMax = elemVal
		}
	}
	return currentMax
}

func max_int_reflect(collection ...interface{}) interface{} {
	currentMax := int64(math.MinInt64)
	for _, elem := range collection {
		elemVal := reflect.ValueOf(elem).Int()
		if elemVal > currentMax {
			currentMax = elemVal
		}
	}
	return currentMax
}

/*
10000000	       129 ns/op	       0 B/op	       0 allocs/op
10000000	       217 ns/op	       8 B/op	       1 allocs/op
2000000	       881 ns/op	       8 B/op	       1 allocs/op
5000000	       261 ns/op	       8 B/op	       1 allocs/op
 */
func Benchmark_int(b *testing.B) {
	fuzzer := fuzz.New()
	datasets := make([][]interface{}, 32)
	typedDatasets := make([][]int, 32)
	for i := 0; i < len(datasets); i++ {
		dataset := make([]interface{}, 100)
		typedDataset := make([]int, 100)
		for j := 0; j < len(dataset); j++ {
			val := int(0)
			fuzzer.Fuzz(&val)
			dataset[j] = val
			typedDataset[j] = val
		}
		datasets[i] = dataset
		typedDatasets[i] = typedDataset
	}
	b.Run("typed", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_int_typed(typedDatasets[i%32]...)
		}
	})
	b.Run("manual", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_int_manual(datasets[i%32]...)
		}
	})
	b.Run("reflect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_int_reflect(datasets[i%32]...)
		}
	})
	b.Run("plz", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			plz.Max(datasets[i%32]...)
		}
	})
}