package acc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

type mapAccessor struct {
	lang.NoopAccessor
	typ                    reflect.Type
	templateEmptyInterface emptyInterface
}

func (accessor *mapAccessor) Kind() lang.Kind {
	return lang.Map
}

func (accessor *mapAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *mapAccessor) Key() lang.Accessor {
	return lang.AccessorOf(reflect.PtrTo(accessor.typ.Key()), accessor.TagName)
}

func (accessor *mapAccessor) Elem() lang.Accessor {
	return lang.AccessorOf(reflect.PtrTo(accessor.typ.Elem()), accessor.TagName)
}

func (accessor *mapAccessor) New() (interface{}, lang.Accessor) {
	return reflect.New(accessor.typ).Interface(), accessor
}

func (accessor *mapAccessor) IterateMap(ptr unsafe.Pointer, cb func(key unsafe.Pointer, value unsafe.Pointer) bool) {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	reflectVal := reflect.ValueOf(castBackEmptyInterface(obj))
	for _, key := range reflectVal.MapKeys() {
		value := reflectVal.MapIndex(key).Interface()
		if !cb(objPtr(key.Interface()), objPtr(value)) {
			return
		}
	}
}

func (accessor *mapAccessor) FillMap(ptr unsafe.Pointer, cb func(filler lang.MapFiller)) {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	filler := &mapFiller{
		typ:   accessor.typ,
		value: reflect.ValueOf(castBackEmptyInterface(obj)),
	}
	cb(filler)
}

type mapFiller struct {
	typ      reflect.Type
	value    reflect.Value
	lastKey  reflect.Value
	lastElem reflect.Value
}

func (filler *mapFiller) Next() (unsafe.Pointer, unsafe.Pointer) {
	filler.lastKey = reflect.New(filler.typ.Key())
	filler.lastElem = reflect.New(filler.typ.Elem())
	return objPtr(filler.lastKey.Interface()), objPtr(filler.lastElem.Interface())
}

func (filler *mapFiller) Fill() {
	filler.value.SetMapIndex(filler.lastKey.Elem(), filler.lastElem.Elem())
}

type mapInterfaceAccessor struct {
	mapAccessor
}

func (accessor *mapInterfaceAccessor) IterateMap(ptr unsafe.Pointer, cb func(key unsafe.Pointer, value unsafe.Pointer) bool) {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	reflectVal := reflect.ValueOf(castBackEmptyInterface(obj))
	for _, key := range reflectVal.MapKeys() {
		value := reflectVal.MapIndex(key).Interface()
		if !cb(objPtr(key.Interface()), objPtr(&value)) {
			return
		}
	}
}
