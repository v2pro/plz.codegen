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

func func_name(typ reflect.Type) string {
	typeName := typ.String()
	typeName = strings.Replace(typeName, ".", "__", -1)
	return typeName
}

func func_symbol(typ reflect.Type) string {
	typeName := func_name(typ)
	typeName = strings.Replace(typeName, "*", "ptr_", -1)
	return typeName
}

func func_is_ptr(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr
}

func func_elem(typ reflect.Type) reflect.Type {
	return typ.Elem()
}

func func_field_of(typ reflect.Type, fieldName string) reflect.StructField {
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic(fieldName + " not found in " + typ.String())
	}
	return field
}

func func_fields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}

func func_is_one_ptr_struct_or_array(typ reflect.Type) bool {
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
