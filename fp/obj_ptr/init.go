package obj_ptr

import (
	"github.com/v2pro/wombat/gen"
	"reflect"
	"unsafe"
)

var F = &gen.FuncTemplate{
	Variables: map[string]string{
		"T": "the type to compare",
	},
	FuncName: `Obj_ptr_{{ .T|symbol }}`,
	Source: `
func {{ .funcName }}(obj interface{}) unsafe.Pointer {
	ptr := (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
	{{ if .T|is_one_ptr_struct_or_array }}
		ptrAsVal := uintptr(ptr)
		ptr = unsafe.Pointer(&ptrAsVal)
	{{ end }}
	return ptr
}
`,
}

func Gen(typ reflect.Type) func(interface{}) unsafe.Pointer {
	funcObj := gen.Compile(F, `T`, typ)
	return funcObj.(func(interface{}) unsafe.Pointer)
}
