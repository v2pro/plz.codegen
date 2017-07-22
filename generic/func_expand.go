package generic

import (
	"text/template"
	"sync"
	"bytes"
	"github.com/v2pro/wombat/compiler"
	"github.com/v2pro/plz"
	"fmt"
	"errors"
)

var logger = plz.LoggerOf("package", "generic")
var expandLock = &sync.Mutex{}
var templates = map[string]*template.Template{}
var state = struct {
	pkgPath           string
	out               *bytes.Buffer
	importPackages    map[string]bool
	declarations      map[string]bool
	expandedFuncNames map[string]bool
}{}
var DynamicCompilationEnabled = false

func Expand(funcTemplate *FuncTemplate, templateArgs ...interface{}) interface{} {
	expandLock.Lock()
	defer expandLock.Unlock()
	state.out = bytes.NewBuffer(nil)
	state.importPackages = map[string]bool{}
	state.declarations = map[string]bool{}
	state.expandedFuncNames = map[string]bool{}
	expandedFuncName, err := funcTemplate.expand(templateArgs)
	if err != nil {
		logger.Error(err, "expand func template failed",
			"funcTemplate", funcTemplate.funcName,
			"templateArgs", templateArgs)
		panic(err.Error())
	}
	expandedFunc := expandedFuncs[expandedFuncName]
	if expandedFunc != nil {
		return expandedFunc
	}
	if !DynamicCompilationEnabled {
		err := logger.Error(nil, "dynamic compilation disabled. " +
			"please add generic.DeclareFunc to init() and re-run codegen",
			"funcTemplate", funcTemplate.funcName,
			"templateArgs", templateArgs,
		"definedInFile", funcTemplate.definedInFile)
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
	for declaration := range state.declarations {
		prelog = prelog + "\n" + declaration
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
	argMap, err := funcTemplate.toArgMap(templateArgs)
	if err != nil {
		return "", err
	}
	localOut := bytes.NewBuffer(nil)
	expandedFuncName := expandSymbolName(funcTemplate.funcName, argMap)
	if state.expandedFuncNames[expandedFuncName] {
		return expandedFuncName, nil
	}
	state.expandedFuncNames[expandedFuncName] = true
	parsedTemplate, err := funcTemplate.parse()
	if err != nil {
		return "", err
	}
	funcTemplate.funcSignature.expand(localOut, expandedFuncName, argMap)
	err = parsedTemplate.Execute(localOut, argMap)
	if err != nil {
		return "", err
	}
	localOut.WriteString("\n}")
	state.out.Write(localOut.Bytes())
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

type ArgMap map[string]interface{}

func (funcTemplate *FuncTemplate) toArgMap(templateArgs []interface{}) (ArgMap, error) {
	argMap := ArgMap{}
	params := map[string]TemplateParam{}
	for k, v := range funcTemplate.templateParams {
		params[k] = v
	}
	for i := 0; i < len(templateArgs); i += 2 {
		argName := templateArgs[i].(string)
		_, found := funcTemplate.templateParams[argName]
		if found {
			delete(params, argName)
		} else {
			return nil, errors.New("argument " + argName + " not declared as param")
		}
		argVal := templateArgs[i+1]
		argMap[argName] = argVal
	}
	for _, param := range params {
		if param.DefaultValueProvider == nil {
			return nil, errors.New("missing param " + param.Name)
		}
		argMap[param.Name] = param.DefaultValueProvider(argMap)
	}
	return argMap, nil
}
