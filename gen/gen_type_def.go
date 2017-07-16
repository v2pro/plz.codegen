package gen

import (
	"bytes"
	"reflect"
	"text/template"
	"strings"
)

func (g *generator) genTypeDef(typ reflect.Type) string {
	if ImportTypes[typ] {
		return ""
	}
	if g.generatedTypes[typ] {
		return ""
	}
	g.generatedTypes[typ] = true
	if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		return g.genTypeDef(typ.Elem())
	}
	if typ.Kind() == reflect.Struct {
		return genStruct(typ)
	}
	if typ.Kind() == reflect.Interface {
		return genInterface(typ)
	}
	return ""

}

func genStruct(typ reflect.Type) string {
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

func genInterface(typ reflect.Type) string {
	if typ.NumMethod() == 0 {
		return ""
	}
	tmpl, err := template.New(typ.String()).Funcs(map[string]interface{}{
		"name":    funcGetName,
		"methods": funcMethods,
	}).Parse(`
type {{ .T|name }} interface {
	{{- range .T|methods }}
	{{ . }}
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

func funcFields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}

func funcMethods(typ reflect.Type) []string {
	methods := []string{}
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		methods = append(methods, strings.Replace(typ.Method(i).Type.String(), "func", method.Name, 1))
	}
	return methods
}
