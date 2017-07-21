package generic

import (
	"reflect"
	"text/template"
	"bytes"
	"strings"
)

func internalizeType(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Struct:
		return internalizeStruct(typ)
	case reflect.Int:
		return "int"
	case reflect.Slice:
		return "[]" + internalizeType(typ.Elem())
	default:
		panic("can not internalize type: " + typ.String())
	}
}

var internalizeStructTmpl *template.Template

func internalizeStruct(typ reflect.Type) string {
	var err error
	if internalizeStructTmpl == nil {
		internalizeStructTmpl, err = template.New(typ.String()).Funcs(map[string]interface{}{
			"name":   genName,
			"fields": genFields,
		}).Parse(`
type {{.structName}} struct {
	{{- range .T|fields }}
	{{.Name}} {{.Type|name}}
	{{- end }}
}`)
		if err != nil {
			panic(err.Error())
		}
	}
	var localOut bytes.Buffer
	structName := strings.Replace(typ.String(), ".", "__", -1)
	err = internalizeStructTmpl.Execute(&localOut, map[string]interface{}{
		"T":          typ,
		"structName": structName,
	})
	if err != nil {
		panic(err.Error())
	}
	state.declarations[localOut.String()] = true
	return structName
}

func genFields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}
