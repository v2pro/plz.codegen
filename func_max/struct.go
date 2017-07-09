package func_max

import (
	"github.com/v2pro/plz/lang"
	"reflect"
	"fmt"
)

func tryMaxStruct(firstElemType reflect.Type, lastElem interface{}) *maxStruct {
	isStruct := firstElemType.Kind() == reflect.Struct && reflect.TypeOf(lastElem).Kind() == reflect.String
	if !isStruct {
		return nil
	}
	sortField := lastElem.(string)
	structAcc := lang.AccessorOf(firstElemType, "max")
	fieldIndex := -1
	var fieldAcc lang.Accessor
	for i := 0; i < structAcc.NumField(); i++ {
		field := structAcc.Field(i)
		if field.Name() == sortField {
			fieldIndex = i
			fieldAcc = field.Accessor()
			break
		}
	}
	if fieldIndex == -1 {
		panic(fmt.Sprintf("sorting field %s can not found int %v", sortField, firstElemType))
	}
	return &maxStruct{
		structAcc:  structAcc,
		fieldIndex: fieldIndex,
		fieldAcc:   fieldAcc,
	}
}

type maxStruct struct {
	structAcc  lang.Accessor
	fieldIndex int
	fieldAcc   lang.Accessor
}

func (f *maxStruct) max(collection []interface{}) interface{} {
	var currentMax interface{}
	currentMaxVal := minInt
	for _, elem := range collection[:len(collection)-1] {
		fieldPtr := f.structAcc.ArrayIndex(lang.AddressOf(elem), f.fieldIndex)
		fieldVal := f.fieldAcc.Int(fieldPtr)
		if fieldVal > currentMaxVal {
			currentMax = elem
			currentMaxVal = fieldVal
		}
	}
	return currentMax
}
