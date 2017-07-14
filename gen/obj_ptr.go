package gen

import (
	"reflect"
	"unsafe"
)

var objPtrF = &FuncTemplate{
	Variables: map[string]string{
		"T": "the type to compare",
	},
	FuncName: `obj_ptr_{{ .T|symbol }}`,
	Source: `
func Exported_{{ .funcName }}(obj interface{}) unsafe.Pointer {
	return {{ .funcName }}(obj)
}
func {{ .funcName }}(obj interface{}) unsafe.Pointer {
	ptr := (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
	{{ if .T|isOnePtrStructOrArray }}
		ptrAsVal := uintptr(ptr)
		ptr = unsafe.Pointer(&ptrAsVal)
	{{ end }}
	return ptr
}
`,
	FuncMap: map[string]interface{}{
		"isOnePtrStructOrArray": funcIsOnePtrStructOrArray,
	},
}

func objPtrGen(typ reflect.Type) func(interface{}) unsafe.Pointer {
	funcObj := Compile(objPtrF, `T`, typ)
	return funcObj.(func(interface{}) unsafe.Pointer)
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
	case reflect.Map:
		return true
	}
	return false
}
