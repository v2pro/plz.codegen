package gen

import (
	"reflect"
	"text/template"
	"bytes"
)

func generateStruct(typ reflect.Type) string {
	if typ.Kind() == reflect.Ptr {
		return generateStruct(typ.Elem())
	}
	if typ.Kind() != reflect.Struct {
		return ""
	}
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