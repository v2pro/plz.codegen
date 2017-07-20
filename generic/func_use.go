package generic

import (
	"reflect"
	"fmt"
	"text/template"
	"sync"
	"bytes"
	"github.com/v2pro/wombat/compiler"
	"github.com/v2pro/plz"
	"strings"
)

var logger = plz.LoggerOf("package", "generic")
var expandLock = &sync.Mutex{}
var templates = map[string]*template.Template{}
var state = struct{
	out *bytes.Buffer
}{}

func Expand(funcTemplate *FuncTemplate, templateArgs ...interface{}) interface{} {
	expandLock.Lock()
	defer expandLock.Unlock()
	state.out = bytes.NewBuffer(nil)
	expandedFuncName, expandedSource, err := funcTemplate.expand(templateArgs)
	if err != nil {
		logger.Error(err, "expand func template failed",
			"funcTemplate", funcTemplate.funcName,
			"templateArgs", templateArgs)
		panic(err.Error())
	}
	prelog := `
package main
	`
	plugin, err := compiler.DynamicCompile(prelog + expandedSource)
	if err != nil {
		panic(err.Error())
	}
	symbol, err := plugin.Lookup(expandedFuncName)
	if err != nil {
		logger.Error(err, "lookup symbol failed",
			"expandedFuncName", expandedFuncName,
			"expandedSource", expandedSource)
		panic(err.Error())
	}
	return symbol
}

func (funcTemplate *FuncTemplate) expand(templateArgs []interface{}) (string, string, error) {
	argMap := map[string]interface{}{}
	for i := 0; i < len(templateArgs); i += 2 {
		argName := templateArgs[i].(string)
		argVal := templateArgs[i+1]
		argMap[argName] = argVal
	}
	expandedFuncName := expandFuncName(funcTemplate.funcName, templateArgs)
	argMap["funcName"] = expandedFuncName
	parsedTemplate, err := funcTemplate.parse()
	if err != nil {
		return "", "", err
	}
	err = parsedTemplate.Execute(state.out, argMap)
	if err != nil {
		return "", "", err
	}
	return expandedFuncName, state.out.String(), nil
}

func (funcTemplate *FuncTemplate) parse() (*template.Template, error) {
	parsedTemplate := templates[funcTemplate.funcName]
	if parsedTemplate == nil {
		var err error
		parsedTemplate, err = template.New(funcTemplate.funcName).
			Funcs(funcTemplate.generators).
			Parse(funcTemplate.templateSource)
		if err != nil {
			return nil, err
		}
		templates[funcTemplate.funcName] = parsedTemplate
	}
	return parsedTemplate, nil
}

func expandFuncName(funcName string, templateArgs []interface{}) string {
	expanded := []byte(funcName)
	for _, arg := range templateArgs {
		switch typedArg := arg.(type) {
		case string:
			expanded = append(expanded, '_')
			expanded = append(expanded, typedArg...)
		case reflect.Type:
			expanded = append(expanded, '_')
			expanded = append(expanded, typeToSymbol(typedArg)...)
		default:
			panic(fmt.Sprintf("unsupported template arg %v of type %s", arg, reflect.TypeOf(arg).String()))
		}
	}
	return string(expanded)
}

func typeToSymbol(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Map:
		return "map_" + typeToSymbol(typ.Key()) + "_to_" + typeToSymbol(typ.Elem())
	case reflect.Slice:
		return "slice_" + typeToSymbol(typ.Elem())
	case reflect.Array:
		return fmt.Sprintf("array_%d_%s", typ.Len(), typeToSymbol(typ.Elem()))
	case reflect.Ptr:
		return "ptr_" + typeToSymbol(typ.Elem())
	default:
		typeName := genName(typ)
		typeName = strings.Replace(typeName, ".", "__", -1)
		if strings.Contains(typeName, "{") {
			typeName = hash(typeName)
		}
		return typeName
	}
}
