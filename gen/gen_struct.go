package gen

import (
	"reflect"
	"text/template"
	"bytes"
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