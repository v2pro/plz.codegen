package fp_compare

import (
	"reflect"
	"html/template"
	"bytes"
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
	if err != nil {
		panic(err.Error())
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]interface{}{
		"T": typ,
	})
	if err != nil {
		panic(err.Error())
	}
	return out.String()
}

func func_name(obj interface{}) string {
	switch typed := obj.(type) {
	case reflect.Type:
		return typed.Name()
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
