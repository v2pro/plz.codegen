package fp_max

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"fmt"
	"unsafe"
)

var cache funcMax

func tryMaxStruct(firstElemType reflect.Type, lastElem interface{}) funcMax {
	isStruct := firstElemType.Kind() == reflect.Struct && reflect.TypeOf(lastElem).Kind() == reflect.String
	if !isStruct {
		return nil
	}
	if cache != nil {
		return cache
	}
	sortField := lastElem.(string)
	var targetField *reflect.StructField
	for i := 0; i < firstElemType.NumField(); i++ {
		field := firstElemType.Field(i)
		if field.Name == sortField {
			targetField = &field
			break
		}
	}
	if targetField == nil {
		panic(fmt.Sprintf("sorting field %s can not found int %v", sortField, firstElemType))
	}
	cache = &funcMaxGeneric{
		&structComparator{
			offset:          targetField.Offset,
			fieldComparator: lang.AccessorOf(intType, "").(lang.Comparator),
		},
	}
	return cache
}

type structComparator struct {
	offset          uintptr
	fieldComparator lang.Comparator
}

func (f *structComparator) Compare(ptr1 unsafe.Pointer, ptr2 unsafe.Pointer) int {
	field1 := unsafe.Pointer(uintptr(ptr1) + f.offset)
	field2 := unsafe.Pointer(uintptr(ptr2) + f.offset)
	return f.fieldComparator.Compare(field1, field2)
}
