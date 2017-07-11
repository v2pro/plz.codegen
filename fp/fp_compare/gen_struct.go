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
	{{ .|name }} {{ .Type|name }}
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

func func_name(obj interface{}) string {
	switch typed := obj.(type) {
	case reflect.Type:
		return strings.Replace(typed.Name(), ".", "__", -1)
	case reflect.StructField:
		return typed.Name
	}
	panic("can not get name from: " + reflect.TypeOf(obj).String())
}

func func_fields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}
