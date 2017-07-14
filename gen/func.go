package gen

import (
	"fmt"
	"reflect"
	"strings"
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func funcGetName(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "int"
	case reflect.Int8:
		return "int8"
	case reflect.Int16:
		return "int16"
	case reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Uint:
		return "uint"
	case reflect.Uint8:
		return "uint8"
	case reflect.Uint16:
		return "uint16"
	case reflect.Uint32:
		return "uint32"
	case reflect.Uint64:
		return "uint64"
	case reflect.Float32:
		return "float32"
	case reflect.Float64:
		return "float64"
	case reflect.String:
		return "string"
	case reflect.Ptr:
		return "*" + funcGetName(typ.Elem())
	}
	typeName := typ.String()
	typeName = strings.Replace(typeName, ".", "__", -1)
	return typeName
}

func funcSymbol(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Map:
		return "map_" + funcSymbol(typ.Key()) + "_to_" + funcSymbol(typ.Elem())
	case reflect.Slice:
		return "slice_" + funcSymbol(typ.Elem())
	case reflect.Array:
		return fmt.Sprintf("array_%d_%s", typ.Len(), funcSymbol(typ.Elem()))
	case reflect.Ptr:
		return "ptr_" + funcSymbol(typ.Elem())
	default:
		typeName := funcGetName(typ)
		if strings.Contains(typeName, "{") {
			typeName = hash(typeName)
		}
		return typeName
	}
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
