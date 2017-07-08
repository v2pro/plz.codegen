package acc

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/lang/tagging"
	"reflect"
	"unsafe"
	"strings"
)

type mapField func(ptr unsafe.Pointer, mapped map[string]interface{})

func accessorOfStruct(typ reflect.Type, tagName string) lang.Accessor {
	tags := tagging.Get(typ)
	fields := []*structField{}
	tagFields := map[string]tagging.FieldTags{}
	for fieldName, fieldTags := range tags.Fields {
		tagFields[fieldName] = fieldTags
	}
	mappedVirtualFields := map[string][]mapField{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldTags := tagFields[field.Name]
		delete(tagFields, field.Name)
		if fieldTags == nil {
			fieldTags = map[string]tagging.TagValue{}
		}
		fieldName := getFieldName(field, fieldTags[tagName])
		if strings.Contains(fieldName, "/") {
			path := strings.Split(fieldName, "/")
			templateEI := castToEmptyInterface(reflect.New(field.Type).Interface())
			lastSection := path[len(path)-1]
			if len(lastSection) > 2 && lastSection[len(lastSection)-2:] == "[]" {
				lastSection = lastSection[:len(lastSection)-2]
				if field.Type.Kind() != reflect.Array && field.Type.Kind() != reflect.Slice {
					arrType := reflect.ArrayOf(1, field.Type)
					templateEI = castToEmptyInterface(reflect.New(arrType).Elem().Interface())
				}
			}
			mappedVirtualFields[path[0]] = append(mappedVirtualFields[path[0]], func(ptr unsafe.Pointer, mapped map[string]interface{}) {
				elemPtr := uintptr(ptr) + field.Offset
				for _, elem := range path[1:len(path)-1] {
					nextLevel := mapped[elem]
					if nextLevel == nil {
						nextLevel = map[string]interface{}{}
						mapped[elem] = nextLevel
					}
					mapped = nextLevel.(map[string]interface{})
				}
				templateEI.word = unsafe.Pointer(elemPtr)
				mapped[lastSection] = castBackEmptyInterface(templateEI)
			})
			continue
		}
		if fieldName == "" {
			continue
		}
		fieldAcc := lang.AccessorOf(reflect.PtrTo(field.Type), tagName)
		fields = append(fields, &structField{
			index:    i,
			name:     fieldName,
			tags:     fieldTags,
			accessor: fieldAcc,
			offset:   field.Offset,
		})
	}
	fields = appendMappedVirtualFields(fields, mappedVirtualFields, tagName)
	fields = appendTagDefinedVirtualFields(fields, tagName, typ, tagFields)
	return &structAccessor{
		NoopAccessor: lang.NoopAccessor{tagName, "structAccessor"},
		typ:          typ,
		fields:       fields,
	}
}

func appendMappedVirtualFields(fields []*structField, mappedVirtualFields map[string][]mapField, tagName string) []*structField {
	index := len(fields)
	for virtualFieldName, mapFields := range mappedVirtualFields {
		fields = append(fields, &structField{
			index:    index,
			name:     virtualFieldName,
			tags:     map[string]tagging.TagValue{},
			accessor: lang.AccessorOf(reflect.TypeOf(map[string]interface{}{}), tagName),
			offset:   0,
			mapValue: func(ptr unsafe.Pointer) interface{} {
				mapped := map[string]interface{}{}
				for _, mapField := range mapFields {
					mapField(ptr, mapped)
				}
				return mapped
			},
		})
		index++
	}
	return fields
}

func appendTagDefinedVirtualFields(fields []*structField, tagName string, typ reflect.Type, tagFields map[string]tagging.FieldTags) []*structField {
	index := len(fields)
	ptr := lang.AddressOf(recursiveNew(typ, 0).Interface())
	for fieldName, fieldTags := range tagFields {
		tagValue := fieldTags[tagName]
		mapValue, _ := tagValue["mapValue"].(func(ptr unsafe.Pointer) interface{})
		if mapValue == nil {
			continue
		}
		value := mapValue(ptr)
		fields = append(fields, &structField{
			index:    index,
			name:     fieldName,
			tags:     fieldTags,
			accessor: lang.AccessorOf(reflect.TypeOf(value), tagName),
			offset:   0,
			mapValue: mapValue,
		})
		index++
	}
	return fields
}

func recursiveNew(typ reflect.Type, level int) reflect.Value {
	pval := reflect.New(typ)
	val := pval.Elem()
	if level < 5 {
		switch typ.Kind() {
		case reflect.Struct:
			for i := 0; i < val.NumField(); i++ {
				field := val.Field(i)
				if field.CanSet() {
					field.Set(recursiveNew(field.Type(), level+1).Elem())
				}
			}
		case reflect.Ptr:
			val.Set(recursiveNew(typ.Elem(), level+1))
		}
	}
	return pval
}

func getFieldName(field reflect.StructField, tagValue tagging.TagValue) string {
	fieldName := field.Name
	fieldTagValue := tagValue.Text()
	if fieldTagValue == "-" {
		fieldName = ""
	} else {
		renameTo := strings.Split(fieldTagValue, ",")[0]
		if renameTo != "" {
			fieldName = renameTo
		}
	}
	return fieldName
}

type structField struct {
	index    int
	name     string
	accessor lang.Accessor
	tags     map[string]tagging.TagValue
	offset   uintptr
	mapValue func(ptr unsafe.Pointer) interface{}
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
	if field.mapValue != nil {
		return objPtr(field.mapValue(ptr))
	}
	return unsafe.Pointer(field.offset + uintptr(ptr))
}

func (accessor *structAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	head := uintptr(ptr)
	for index := 0; index < len(accessor.fields); index++ {
		field := accessor.fields[index]
		elemPtr := unsafe.Pointer(head + field.offset)
		if field.mapValue != nil {
			elemPtr = objPtr(field.mapValue(ptr))
		}
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
