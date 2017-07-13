package gen

import (
	"bytes"
	"reflect"
	"text/template"
)

func (g *generator) genStruct(typ reflect.Type) string {
	if g.generatedTypes[typ] {
		return ""
	}
	g.generatedTypes[typ] = true
	if typ.Kind() == reflect.Ptr {
		return g.genStruct(typ.Elem())
	}
	if typ.Kind() != reflect.Struct {
		return ""
	}
	tmpl, err := template.New(typ.String()).Funcs(map[string]interface{}{
		"name":   funcGetName,
		"fields": funcFields,
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
