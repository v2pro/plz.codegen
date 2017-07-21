package max

import (
	"testing"
	"github.com/v2pro/wombat/generic"
	"reflect"
	"github.com/stretchr/testify/require"
	"github.com/google/gofuzz"
)

func Test_slice_int(t *testing.T) {
	should := require.New(t)
	f := generic.Expand(ByItself, "T", reflect.TypeOf([]int{})).
	(func([]int) int)
	should.Equal(3, f([]int{1, 3, 2}))
}

func Benchmark_slice_int(b *testing.B) {
	fuzzer := fuzz.New()
	datasets := make([][]int, 32)
	typedDatasets := make([][]int, 32)
	for i := 0; i < len(datasets); i++ {
		dataset := make([]int, 100)
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
	f := generic.Expand(ByItself, "T", reflect.TypeOf([]int{})).
	(func([]int) int)
	b.Run("plz", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			f(datasets[i%32])
		}
	})
}