package cp_json

import (
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/lang"
	"unsafe"
)

type streamAccessor struct {
	lang.NoopAccessor
	kind lang.Kind
}

func (accessor *streamAccessor) ReadOnly() bool {
	return false
}

func (accessor *streamAccessor) Kind() lang.Kind {
	return accessor.kind
}

func (accessor *streamAccessor) RandomAccessible() bool {
	return false
}

func (accessor *streamAccessor) GoString() string {
	return "cp_json.streamAccessor"
}

func (accessor *streamAccessor) Key() lang.Accessor {
	return &mapKeyWriter{}
}

func (accessor *streamAccessor) Elem() lang.Accessor {
	return &streamAccessor{
		lang.NoopAccessor{accessor.TagName, "streamAccessor"},
		lang.Variant,
	}
}

func (accessor *streamAccessor) Skip(ptr unsafe.Pointer) {
	stream := (*jsoniter.Stream)(ptr)
	stream.WriteNil()
}

func (accessor *streamAccessor) SetString(ptr unsafe.Pointer, val string) {
	stream := (*jsoniter.Stream)(ptr)
	stream.WriteString(val)
}

func (accessor *streamAccessor) SetInt(ptr unsafe.Pointer, val int) {
	stream := (*jsoniter.Stream)(ptr)
	stream.WriteInt(val)
}

func (accessor *streamAccessor) FillMap(ptr unsafe.Pointer, cb func(filler lang.MapFiller)) {
	stream := (*jsoniter.Stream)(ptr)
	stream.WriteObjectStart()
	cb(&jsonMapFiller{stream, true})
	stream.WriteObjectEnd()
}

type jsonMapFiller struct {
	stream  *jsoniter.Stream
	isFirst bool
}

func (filler *jsonMapFiller) Next() (unsafe.Pointer, unsafe.Pointer) {
	if filler.isFirst {
		filler.isFirst = false
	} else {
		filler.stream.WriteMore()
	}
	return unsafe.Pointer(filler.stream), unsafe.Pointer(filler.stream)
}

func (filler *jsonMapFiller) Fill() {
}

func (accessor *streamAccessor) FillArray(ptr unsafe.Pointer, cb func(filler lang.ArrayFiller)) {
	stream := (*jsoniter.Stream)(ptr)
	stream.WriteArrayStart()
	cb(&jsonArrayFiller{
		stream:  stream,
		isFirst: true,
	})
	stream.WriteArrayEnd()
}

type jsonArrayFiller struct {
	stream  *jsoniter.Stream
	index   int
	isFirst bool
}

func (filler *jsonArrayFiller) Next() (int, unsafe.Pointer) {
	if filler.isFirst {
		filler.isFirst = false
	} else {
		filler.stream.WriteMore()
	}
	currentIndex := filler.index
	filler.index++
	return currentIndex, unsafe.Pointer(filler.stream)
}

func (filler *jsonArrayFiller) Fill() {
}

type mapKeyWriter struct {
	lang.NoopAccessor
}

func (accessor *mapKeyWriter) ReadOnly() bool {
	return false
}

func (accessor *mapKeyWriter) Kind() lang.Kind {
	return lang.String
}

func (accessor *mapKeyWriter) GoString() string {
	return "string"
}

func (accessor *mapKeyWriter) SetString(ptr unsafe.Pointer, val string) {
	stream := (*jsoniter.Stream)(ptr)
	stream.WriteObjectField(val)
}
