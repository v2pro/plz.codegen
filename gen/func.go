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
