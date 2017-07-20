package generic

import (
	"text/template"
	"sync"
	"bytes"
	"github.com/v2pro/wombat/compiler"
	"github.com/v2pro/plz"
	"fmt"
)

var logger = plz.LoggerOf("package", "generic")
var expandLock = &sync.Mutex{}
var templates = map[string]*template.Template{}
var state = struct{
	out *bytes.Buffer
	importPackages map[string]bool
}{}

func Expand(funcTemplate *FuncTemplate, templateArgs ...interface{}) interface{} {
	expandLock.Lock()
	defer expandLock.Unlock()
	state.out = bytes.NewBuffer(nil)
	state.importPackages = map[string]bool{}
	expandedFuncName, err := funcTemplate.expand(templateArgs)
	if err != nil {
		logger.Error(err, "expand func template failed",
			"funcTemplate", funcTemplate.funcName,
			"templateArgs", templateArgs)
		panic(err.Error())
	}
	prelog := `
package main
	`
	for importPackage := range state.importPackages {
		prelog = fmt.Sprintf(`
%s
import "%s"`, prelog, importPackage)
	}
	expandedSource := state.out.String()
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

func (funcTemplate *FuncTemplate) expand(templateArgs []interface{}) (string, error) {
	argMap := map[string]interface{}{}
	for i := 0; i < len(templateArgs); i += 2 {
		argName := templateArgs[i].(string)
		argVal := templateArgs[i+1]
		argMap[argName] = argVal
	}
	expandedFuncName := expandSymbolName(funcTemplate.funcName, templateArgs)
	argMap["funcName"] = expandedFuncName
	parsedTemplate, err := funcTemplate.parse()
	if err != nil {
		return "", err
	}
	err = parsedTemplate.Execute(state.out, argMap)
	if err != nil {
		return "", err
	}
	return expandedFuncName, nil
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
