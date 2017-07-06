package jsonacc
//
//import (
//	"testing"
//	"github.com/json-iterator/go"
//	"github.com/v2pro/plz"
//	"reflect"
//	"github.com/stretchr/testify/require"
//	"github.com/v2pro/plz/lang"
//	"unsafe"
//)
//
//func Test_map_decode(t *testing.T) {
//	should := require.New(t)
//	iter := jsoniter.ParseString(jsoniter.ConfigDefault,
//		`{"Field": 1}`)
//	accessor := lang.AccessorOf(reflect.TypeOf(iter))
//	should.Equal(lang.Variant, accessor.Kind())
//	should.Equal(lang.String, accessor.Key().Kind())
//	should.Equal(lang.Variant, accessor.Elem().Kind())
//	keys := []string{}
//	elems := []int{}
//	accessor.IterateMap(iter, func(key unsafe.Pointer, elem unsafe.Pointer) bool {
//		keys = append(keys, accessor.Key().String(key))
//		elems = append(elems, accessor.Elem().Int(elem))
//		return true
//	})
//	should.Equal([]string{"Field "}, keys)
//	should.Equal([]int{1}, elems)
//}
//
//func Test_map_encode_one(t *testing.T) {
//	should := require.New(t)
//	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
//	accessor := lang.AccessorOf(reflect.TypeOf(stream))
//	accessor.FillMap(stream, func(filler lang.MapFiller) {
//		key, elem := filler.Next()
//		accessor.Key().SetString(key, "hello")
//		accessor.Elem().SetString(elem, "world")
//		filler.Fill()
//	})
//	should.Equal(`{"hello":"world"}`, string(stream.Buffer()))
//}
//
//func Test_map_encode_many(t *testing.T) {
//	should := require.New(t)
//	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 1024)
//	accessor := lang.AccessorOf(reflect.TypeOf(stream))
//	accessor.FillMap(stream, func(filler lang.MapFiller) {
//		key, elem := filler.Next()
//		accessor.Key().SetString(key, "hello")
//		accessor.Elem().SetString(elem, "world")
//		filler.Fill()
//		key, elem = filler.Next()
//		accessor.Key().SetString(key, "k")
//		accessor.Elem().SetString(elem, "v")
//		filler.Fill()
//	})
//	should.Equal(`{"hello":"world","k":"v"}`, string(stream.Buffer()))
//}
