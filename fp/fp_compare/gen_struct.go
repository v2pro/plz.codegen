package fp_compare

import (
	"reflect"
	"text/template"
	"bytes"
	"strings"
)

func generateStruct(typ reflect.Type) string {
	tmpl, err := template.New(typ.String()).Funcs(map[string]interface{}{
		"name":   func_name,
		"fields": func_fields,
	}).Parse(`
type {{ .T|name }} struct {
	{{- range .T|fields }}
	{{ .Name }} {{ .Type|name }}
	{{- end }}
}`)
	panicOnError(err)
	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]interface{}{
		"T": typ,
	})
	panicOnError(err)
	return out.String()
}

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func func_name(typ reflect.Type) string {
	return strings.Replace(typ.String(), ".", "__", -1)
}

func func_fields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}

const genObjPtr = `
func objPtr(obj interface{}) unsafe.Pointer {
	emptyInterface := *((*emptyInterface)(unsafe.Pointer(&obj)))
	ptr := emptyInterface.word
	switch reflect.Kind(emptyInterface.typ.kind & kindMask) {
	case reflect.Array:
		typ := reflect.TypeOf(obj)
		if typ.Len() == 1 && (typ.Elem().Kind() == reflect.Ptr || typ.Elem().Kind() == reflect.Map) {
			asVal := uintptr(ptr)
			ptr = unsafe.Pointer(&asVal)
		}
	case reflect.Struct:
		typ := reflect.TypeOf(obj)
		if typ.NumField() == 1 && (typ.Field(0).Type.Kind() == reflect.Ptr || typ.Field(0).Type.Kind() == reflect.Map) {
			asVal := uintptr(ptr)
			ptr = unsafe.Pointer(&asVal)
		}
	}
	return ptr
}

func castToEmptyInterface(obj interface{}) emptyInterface {
	return *((*emptyInterface)(unsafe.Pointer(&obj)))
}

const kindMask = (1 << 5) - 1

type rtype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32 // hash of type; avoids computation in hash tables
	tflag      uint8  // extra type information flags
	align      uint8  // alignment of variable with this type
	fieldAlign uint8  // alignment of struct field with this type
	kind       uint8  // enumeration for C
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}
`

func func_cast(identifier string, typ reflect.Type) string {
	return "*(*" + func_name(typ) + ")(objPtr(" + identifier + "))"
}