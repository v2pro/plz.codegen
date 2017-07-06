package jsonacc

import (
	"github.com/v2pro/plz/lang"
	"github.com/json-iterator/go"
)

type streamAccessor struct {
	lang.NoopAccessor
}

func (accessor *streamAccessor) Kind() lang.Kind {
	return lang.Variant
}

func (accessor *streamAccessor) GoString() string {
	return "interface{}"
}

func (accessor *streamAccessor) Key() lang.Accessor {
	return &mapKeyWriter{}
}

func (accessor *streamAccessor) Elem() lang.Accessor {
	return accessor
}

func (accessor *streamAccessor) SetString(obj interface{}, val string) {
	stream := obj.(*jsoniter.Stream)
	stream.WriteString(val)
}

func (accessor *streamAccessor) SetInt(obj interface{}, val int) {
	stream := obj.(*jsoniter.Stream)
	stream.WriteInt(val)
}

func (accessor *streamAccessor) FillMap(obj interface{}, cb func(filler lang.MapFiller)) {
	stream := obj.(*jsoniter.Stream)
	stream.WriteObjectStart()
	cb(&jsonMapFiller{stream, obj, true})
	stream.WriteObjectEnd()
}

type jsonMapFiller struct {
	stream  *jsoniter.Stream
	obj     interface{}
	isFirst bool
}

func (filler *jsonMapFiller) Next() (interface{}, interface{}) {
	if filler.isFirst {
		filler.isFirst = false
	} else {
		filler.stream.WriteMore()
	}
	return filler.obj, filler.obj
}

func (filler *jsonMapFiller) Fill() {
}

func (accessor *streamAccessor) FillArray(obj interface{}, cb func(filler lang.ArrayFiller)) {
	stream := obj.(*jsoniter.Stream)
	stream.WriteArrayStart()
	cb(&jsonArrayFiller{
		stream:  stream,
		obj:     obj,
		isFirst: true,
	})
	stream.WriteArrayEnd()
}

type jsonArrayFiller struct {
	stream  *jsoniter.Stream
	index   int
	obj     interface{}
	isFirst bool
}

func (filler *jsonArrayFiller) Next() (int, interface{}) {
	if filler.isFirst {
		filler.isFirst = false
	} else {
		filler.stream.WriteMore()
	}
	currentIndex := filler.index
	filler.index++
	return currentIndex, filler.obj
}

func (filler *jsonArrayFiller) Fill() {
}

type mapKeyWriter struct {
	lang.NoopAccessor
}

func (accessor *mapKeyWriter) Kind() lang.Kind {
	return lang.String
}

func (accessor *mapKeyWriter) GoString() string {
	return "string"
}

func (accessor *mapKeyWriter) SetString(obj interface{}, val string) {
	stream := obj.(*jsoniter.Stream)
	stream.WriteObjectField(val)
}
