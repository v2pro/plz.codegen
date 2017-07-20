package generic

import (
	"reflect"
	"fmt"
	"strings"
)

func expandSymbolName(plainName string, templateArgs []interface{}) string {
	expanded := []byte(plainName)
	for _, arg := range templateArgs {
		switch typedArg := arg.(type) {
		case string:
			expanded = append(expanded, '_')
			expanded = append(expanded, typedArg...)
		case reflect.Type:
			expanded = append(expanded, '_')
			expanded = append(expanded, typeToSymbol(typedArg)...)
		default:
			panic(fmt.Sprintf("unsupported template arg %v of type %s", arg, reflect.TypeOf(arg).String()))
		}
	}
	return string(expanded)
}

func typeToSymbol(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Map:
		return "map_" + typeToSymbol(typ.Key()) + "_to_" + typeToSymbol(typ.Elem())
	case reflect.Slice:
		return "slice_" + typeToSymbol(typ.Elem())
	case reflect.Array:
		return fmt.Sprintf("array_%d_%s", typ.Len(), typeToSymbol(typ.Elem()))
	case reflect.Ptr:
		return "ptr_" + typeToSymbol(typ.Elem())
	default:
		typeName := typ.String()
		typeName = strings.Replace(typeName, ".", "__", -1)
		if strings.Contains(typeName, "{") {
			typeName = hash(typeName)
		}
		return typeName
	}
}
