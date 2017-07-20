package generic

import (
	"reflect"
	"fmt"
	"text/template"
	"sync"
	"bytes"
	"github.com/v2pro/wombat/compiler"
	"github.com/v2pro/plz"
)

var logger = plz.LoggerOf("package", "generic")
var expandLock = &sync.Mutex{}
var templates = map[string]*template.Template{}

func Expand(funcTemplate *FuncTemplate, templateArgs ...interface{}) interface{} {
	expandLock.Lock()
	defer expandLock.Unlock()
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
	out := bytes.NewBuffer(nil)
	err = parsedTemplate.Execute(out, argMap)
	if err != nil {
		return "", "", err
	}
	return expandedFuncName, out.String(), nil
}

func (funcTemplate *FuncTemplate) parse() (*template.Template, error) {
	parsedTemplate := templates[funcTemplate.funcName]
	if parsedTemplate == nil {
		genMap := template.FuncMap{
			"name": genName,
		}
		var err error
		parsedTemplate, err = template.New(funcTemplate.funcName).
			Funcs(genMap).
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
	return typ.String()
}
