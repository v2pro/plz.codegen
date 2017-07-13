package gen

import (
	"reflect"
	"strings"
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func funcGetName(typ reflect.Type) string {
	typeName := typ.String()
	typeName = strings.Replace(typeName, ".", "__", -1)
	return typeName
}

func funcSymbol(typ reflect.Type) string {
	typeName := funcGetName(typ)
	typeName = strings.Replace(typeName, "*", "ptr_", -1)
	typeName = strings.Replace(typeName, "[]", "slice_", -1)
	return typeName
}

func funcIsPtr(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr
}

func funcElem(typ reflect.Type) reflect.Type {
	return typ.Elem()
}

func funcFieldOf(typ reflect.Type, fieldName string) reflect.StructField {
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic(fieldName + " not found in " + typ.String())
	}
	return field
}

func funcFields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}

func funcIsOnePtrStructOrArray(typ reflect.Type) bool {
	switch reflect.Kind(typ.Kind()) {
	case reflect.Array:
		if typ.Len() == 1 && (typ.Elem().Kind() == reflect.Ptr || typ.Elem().Kind() == reflect.Map) {
			return true
		}
	case reflect.Struct:
		if typ.NumField() == 1 && (typ.Field(0).Type.Kind() == reflect.Ptr || typ.Field(0).Type.Kind() == reflect.Map) {
			return true
		}
	}
	return false
}
