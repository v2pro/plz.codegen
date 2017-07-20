package generic

import (
	"text/template"
	"reflect"
	"fmt"
)

type structCacheKey struct {
	structName    string
	interfaceType reflect.Type
}

var structCtors = map[structCacheKey]func() interface{}{
}

func New(structTemplate *StructTemplate, interfaceType reflect.Type) interface{} {
	cacheKey := structCacheKey{
		structName:    structTemplate.structName,
		interfaceType: interfaceType,
	}
	ctor := structCtors[cacheKey]
	if ctor == nil {
		ctor = structTemplate.expandCtor(interfaceType)
		structCtors[cacheKey] = ctor
	}
	return ctor()
}

func (structTemplate *StructTemplate) expandCtor(interfaceType reflect.Type) func() interface{} {
	f := Func("New_" + structTemplate.structName).
		Params("I", "interface of the expanded struct").
		ImportStruct(structTemplate).
		Source(fmt.Sprintf(`
{{ $struct := expand "%s" "I" .I }}

func {{.funcName}} () interface{} {
    return &{{$struct}}{}
}`, structTemplate.structName))
	ctor := Expand(f, "I", interfaceType).(func() interface{})
	return ctor
}

func (structTemplate *StructTemplate) expand(templateArgs []interface{}) (string, error) {
	argMap := map[string]interface{}{}
	for i := 0; i < len(templateArgs); i += 2 {
		argName := templateArgs[i].(string)
		argVal := templateArgs[i+1]
		argMap[argName] = argVal
	}
	expandedStructName := expandSymbolName(structTemplate.structName, templateArgs)
	argMap["structName"] = expandedStructName
	parsedTemplate, err := structTemplate.parse()
	if err != nil {
		return "", err
	}
	err = parsedTemplate.Execute(state.out, argMap)
	if err != nil {
		return "", err
	}
	return expandedStructName, nil
}

func (structTemplate *StructTemplate) parse() (*template.Template, error) {
	parsedTemplate := templates[structTemplate.structName]
	if parsedTemplate == nil {
		var err error
		parsedTemplate, err = template.New(structTemplate.structName).
			Funcs(structTemplate.generators).
			Parse(structTemplate.templateSource)
		if err != nil {
			return nil, err
		}
		templates[structTemplate.structName] = parsedTemplate
	}
	return parsedTemplate, nil
}
