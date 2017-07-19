package generic

import (
	"reflect"
	"fmt"
	"text/template"
	"sync"
	"bytes"
)

var expandLock = &sync.Mutex{}
var templates = map[string]*template.Template{}

func Expand(funcTemplate *FuncTemplate, templateArgs ...interface{}) interface{} {
	expandLock.Lock()
	defer expandLock.Unlock()
}

func (funcTemplate *FuncTemplate) render(templateArgs []interface{}) (string, error) {
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
		return "", err
	}
	var out bytes.Buffer
	err = parsedTemplate.Execute(out, argMap)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (funcTemplate *FuncTemplate) parse() (*template.Template, error) {
	parsedTemplate := templates[funcTemplate.funcName]
	if parsedTemplate == nil {
		var err error
		parsedTemplate, err = template.New(funcTemplate.funcName).
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
