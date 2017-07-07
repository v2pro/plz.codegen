package acc

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/lang/tagging"
	"reflect"
	"unsafe"
	"strings"
)

func accessorOfStruct(typ reflect.Type, tagName string) lang.Accessor {
	tags := tagging.Get(typ)
	fields := []*structField{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldAcc := lang.AccessorOf(reflect.PtrTo(field.Type), tagName)
		fieldTags := tags.Fields[field.Name]
		if fieldTags == nil {
			fieldTags = map[string]tagging.TagValue{}
		}
		fieldName := field.Name
		fieldTagValue := fieldTags[tagName].Text()
		if fieldTagValue == "-" {
			fieldName = ""
		} else {
			renameTo := strings.Split(fieldTagValue, ",")[0]
			if renameTo != "" {
				fieldName = renameTo
			}
		}
		fields = append(fields, &structField{
			index:    i,
			name:     fieldName,
			tags:     fieldTags,
			accessor: fieldAcc,
			offset:   field.Offset,
		})
	}
	return &structAccessor{
		NoopAccessor: lang.NoopAccessor{tagName,"structAccessor"},
		typ:          typ,
		fields:       fields,
	}
}

type structField struct {
	index    int
	name     string
	accessor lang.Accessor
	tags     map[string]tagging.TagValue
	offset   uintptr
}

func (sf *structField) Index() int {
	return sf.index
}

func (sf *structField) Name() string {
	return sf.name
}

func (sf *structField) Accessor() lang.Accessor {
	return sf.accessor
}

func (sf *structField) Tags() map[string]tagging.TagValue {
	return sf.tags
}

type structAccessor struct {
	lang.NoopAccessor
	typ    reflect.Type
	fields []*structField
}

func (accessor *structAccessor) Kind() lang.Kind {
	return lang.Struct
}

func (accessor *structAccessor) GoString() string {
	return accessor.typ.String()
}

func (accessor *structAccessor) NumField() int {
	return len(accessor.fields)
}

func (accessor *structAccessor) Field(index int) lang.StructField {
	return accessor.fields[index]
}

func (accessor *structAccessor) RandomAccessible() bool {
	return true
}

func (accessor *structAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	field := accessor.fields[index]
	return unsafe.Pointer(field.offset + uintptr(ptr))
}

func (accessor *structAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	head := uintptr(ptr)
	for index := 0; index < len(accessor.fields); index++ {
		field := accessor.fields[index]
		elemPtr := unsafe.Pointer(head + field.offset)
		if field.accessor.IsNil(elemPtr) {
			elemPtr = nil
		}
		cb(index, elemPtr)
	}
}

func (accessor *structAccessor) New() (interface{}, lang.Accessor) {
	ptrAcc := lang.AccessorOf(reflect.PtrTo(accessor.typ), accessor.TagName)
	return reflect.New(accessor.typ).Elem().Interface(), ptrAcc
}

type ptrStructAccessor struct {
	ptrAccessor
}

func (accessor *ptrStructAccessor) ArrayIndex(ptr unsafe.Pointer, index int) unsafe.Pointer {
	return accessor.valueAccessor.ArrayIndex(ptr, index)
}

func (accessor *ptrStructAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	accessor.valueAccessor.IterateArray(ptr, cb)
}

func (accessor *ptrStructAccessor) FillArray(ptr unsafe.Pointer, cb func(filler lang.ArrayFiller)) {
	fields := accessor.valueAccessor.(*structAccessor).fields
	filler := &structFiller{
		fields: fields,
		head:   uintptr(ptr),
	}
	cb(filler)
}

func (accessor *ptrStructAccessor) New() (interface{}, lang.Accessor) {
	typ := accessor.valueAccessor.(*structAccessor).typ
	return reflect.New(typ).Elem().Interface(), accessor
}

type structFiller struct {
	fields []*structField
	index  int
	head   uintptr
}

func (filler *structFiller) Next() (int, unsafe.Pointer) {
	field := filler.fields[filler.index]
	currentIndex := filler.index
	filler.index++
	return currentIndex, unsafe.Pointer(filler.head + field.offset)
}

func (filler *structFiller) Fill() {
}
