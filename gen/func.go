package gen

import (
	"fmt"
	"reflect"
	"strings"
)

var ImportPackages = map[string]bool {
}

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func funcGetName(typ reflect.Type) string {
	if ImportPackages[typ.PkgPath()] {
		return typ.String()
	}
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
	case reflect.Slice:
		return "[]" + funcGetName(typ.Elem())
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", typ.Len(), funcGetName(typ.Elem()))
	case reflect.Map:
		return "map[" + funcGetName(typ.Key()) + "]" + funcGetName(typ.Elem())
	}
	typeName := typ.String()
	typeName = strings.Replace(typeName, ".", "__", -1)
	if strings.Contains(typeName, "struct {") {
		typeName = hash(typeName)
	}
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
		typeName = strings.Replace(typeName, ".", "__", -1)
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
	switch typ.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.Ptr:
		return typ.Elem()
	}
	panic("can not get elem from " + typ.String())
}

func funcFieldOf(typ reflect.Type, fieldName string) reflect.StructField {
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic(fieldName + " not found in " + typ.String())
	}
	return field
}
