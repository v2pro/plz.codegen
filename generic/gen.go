package generic

import (
	"reflect"
	"fmt"
	"strings"
	"crypto/sha1"
	"encoding/base32"
	"runtime"
)

func genName(typ reflect.Type) string {
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
	case reflect.Bool:
		return "bool"
	case reflect.Ptr:
		return "*" + genName(typ.Elem())
	case reflect.Slice:
		return "[]" + genName(typ.Elem())
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", typ.Len(), genName(typ.Elem()))
	case reflect.Map:
		return "map[" + genName(typ.Key()) + "]" + genName(typ.Elem())
	case reflect.Struct, reflect.Interface:
		typeName := typ.String()
		typeName = strings.Replace(typeName, ".", "__", -1)
		if strings.Contains(typeName, "struct {") {
			typeName = hash(typeName)
		}
		return typeName
	}
	panic("do not support " + typ.Kind().String())
}

func genElem(typ reflect.Type) reflect.Type {
	return typ.Elem()
}

func hash(source string) string {
	h := sha1.New()
	h.Write([]byte(source))
	h.Write([]byte(runtime.Version()))
	return "g" + base32.StdEncoding.EncodeToString(h.Sum(nil))
}