package acc

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

func accessorOfMap(typ reflect.Type, tagName string) lang.Accessor {
	templateEmptyInterface := castToEmptyInterface(reflect.New(typ).Elem().Interface())
	templateKeyEmptyInterface := castToEmptyInterface(reflect.New(typ.Key()).Interface())
	templateElemEmptyInterface := castToEmptyInterface(reflect.New(typ.Elem()).Interface())
	if typ.Elem().Kind() == reflect.Interface {
		return &mapInterfaceAccessor{
			mapAccessor{
				NoopAccessor:               lang.NoopAccessor{tagName, "mapInterfaceAccessor"},
				typ:                        typ,
				templateEmptyInterface:     templateEmptyInterface,
				templateKeyEmptyInterface:  templateKeyEmptyInterface,
				templateElemEmptyInterface: templateElemEmptyInterface,
			},
		}
	}
	return &mapAccessor{
		NoopAccessor:               lang.NoopAccessor{tagName, "mapAccessor"},
		typ:                        typ,
		templateEmptyInterface:     templateEmptyInterface,
		templateKeyEmptyInterface:  templateKeyEmptyInterface,
		templateElemEmptyInterface: templateElemEmptyInterface,
	}
}

type mapAccessor struct {
	lang.NoopAccessor
	typ                        reflect.Type
	templateEmptyInterface     emptyInterface
	templateKeyEmptyInterface  emptyInterface
	templateElemEmptyInterface emptyInterface
}

func (accessor *mapAccessor) Kind() lang.Kind {
	return lang.Map
}

func (accessor *mapAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *mapAccessor) RandomAccessible() bool {
	return true
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

func (accessor *mapAccessor) MapIndex(ptr unsafe.Pointer, keyPtr unsafe.Pointer) unsafe.Pointer {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	reflectVal := reflect.ValueOf(castBackEmptyInterface(obj))
	key := accessor.templateKeyEmptyInterface
	key.word = keyPtr
	value := reflectVal.MapIndex(reflect.ValueOf(castBackEmptyInterface(key)).Elem())
	if !value.IsValid() {
		return unsafe.Pointer(nil)
	}
	if accessor.typ.Elem().Kind() == reflect.Ptr || accessor.typ.Elem().Kind() == reflect.Interface {
		elemPtr := uintptr(objPtr(value.Interface()))
		if elemPtr != 0 {
			return unsafe.Pointer(&elemPtr)
		} else {
			return unsafe.Pointer(nil)
		}
	} else {
		return objPtr(value.Interface())
	}
}

func (accessor *mapAccessor) SetMapIndex(ptr unsafe.Pointer, keyPtr unsafe.Pointer, elemPtr unsafe.Pointer) {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	reflectVal := reflect.ValueOf(castBackEmptyInterface(obj))
	key := accessor.templateKeyEmptyInterface
	key.word = keyPtr
	elem := accessor.templateElemEmptyInterface
	elem.word = elemPtr
	reflectVal.SetMapIndex(
		reflect.ValueOf(castBackEmptyInterface(key)).Elem(),
		reflect.ValueOf(castBackEmptyInterface(elem)).Elem())
}

func (accessor *mapAccessor) IterateMap(ptr unsafe.Pointer, cb func(key unsafe.Pointer, value unsafe.Pointer) bool) {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	reflectVal := reflect.ValueOf(castBackEmptyInterface(obj))
	for _, key := range reflectVal.MapKeys() {
		keyPtr := objPtr(key.Interface())
		elem := reflectVal.MapIndex(key).Interface()
		elemPtr := objPtr(elem)
		if accessor.typ.Elem().Kind() == reflect.Ptr || accessor.typ.Elem().Kind() == reflect.Interface {
			asValue := uintptr(elemPtr)
			if asValue != 0 {
				elemPtr = unsafe.Pointer(&asValue)
			}
		}
		if !cb(keyPtr, elemPtr) {
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
		value := reflectVal.MapIndex(key)
		if value.IsNil() {
			if !cb(objPtr(key.Interface()), nil) {
				return
			}
		} else {
			valueAsInterface := value.Interface()
			if !cb(objPtr(key.Interface()), objPtr(&valueAsInterface)) {
				return
			}
		}
	}
}

func (accessor *mapInterfaceAccessor) MapIndex(ptr unsafe.Pointer, keyPtr unsafe.Pointer) unsafe.Pointer {
	obj := accessor.templateEmptyInterface
	obj.word = ptr
	reflectVal := reflect.ValueOf(castBackEmptyInterface(obj))
	key := accessor.templateKeyEmptyInterface
	key.word = keyPtr
	value := reflectVal.MapIndex(reflect.ValueOf(castBackEmptyInterface(key)).Elem())
	if !value.IsValid() || value.IsNil() {
		return unsafe.Pointer(nil)
	}
	valueAsInterface := value.Interface()
	return objPtr(&valueAsInterface)
}
