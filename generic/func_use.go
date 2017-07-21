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
	testMode bool
	out *bytes.Buffer
	importPackages map[string]bool
	declarations map[string]bool
}{}

func Expand(funcTemplate *FuncTemplate, templateArgs ...interface{}) interface{} {
	expandLock.Lock()
	defer expandLock.Unlock()
	state.out = bytes.NewBuffer(nil)
	state.importPackages = map[string]bool{}
	state.declarations = map[string]bool{}
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
	argMap := map[string]interface{}{}
	for i := 0; i < len(templateArgs); i += 2 {
		argName := templateArgs[i].(string)
		argVal := templateArgs[i+1]
		argMap[argName] = argVal
	}
	localOut := bytes.NewBuffer(nil)
	expandedFuncName := expandSymbolName(funcTemplate.funcName, argMap)
	testMode := argMap["testMode"] == true
	if testMode {
		err := funcTemplate.expandTestModeEntryFunc(localOut, expandedFuncName, argMap)
		if err != nil {
			return "", err
		}
	} else {
		parsedTemplate, err := funcTemplate.parse()
		if err != nil {
			return "", err
		}
		err = funcTemplate.funcSignature.expand(localOut, expandedFuncName, argMap)
		if err != nil {
			return "", err
		}
		err = parsedTemplate.Execute(localOut, argMap)
		if err != nil {
			return "", err
		}
		localOut.WriteString("\n}")
	}
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
