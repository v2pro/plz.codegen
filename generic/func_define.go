package generic

import (
	"fmt"
	"reflect"
	"runtime"
)

type FuncTemplateBuilder struct {
	funcTemplate *FuncTemplate
}

func DefineFunc(signature string) *FuncTemplateBuilder {
	importedFuncTemplates := map[string]*FuncTemplate{}
	importedStructTemplates := map[string]*StructTemplate{}
	parsedSignature, err := parseSignature(signature)
	if err != nil {
		panic(err.Error())
	}
	_, definedInFile, _, _ := runtime.Caller(1)
	return &FuncTemplateBuilder{funcTemplate: &FuncTemplate{
		funcSignature:           parsedSignature,
		definedInFile:           definedInFile,
		templateParams:          map[string]TemplateParam{},
		importedFuncTemplates:   importedFuncTemplates,
		importedStructTemplates: importedStructTemplates,
		importedPackages: map[string]bool{},
		generators: map[string]interface{}{
			"name": genName,
			"elem": genElem,
			"expand": func(depName string, templateArgs ...interface{}) (string, error) {
				depFunc := importedFuncTemplates[depName]
				if depFunc != nil {
					expandedFuncName, err := depFunc.expand(templateArgs)
					if err != nil {
						return "", logger.Error(err, fmt.Sprintf("expand func %s failed", depName),
							"templateArgs", templateArgs)
					}
					return expandedFuncName, nil
				}
				depStruct := importedStructTemplates[depName]
				if depStruct != nil {
					expandedStructName, err := depStruct.expand(templateArgs)
					if err != nil {
						return "", logger.Error(err, fmt.Sprintf("expand struct %s failed", depName),
							"templateArgs", templateArgs)
					}
					return expandedStructName, nil
				}
				logger.Error(nil, "missing dependency", "depName", depName)
				return "", fmt.Errorf(
					"referenced generic function %s should be imported by ImportFunc",
					depName)
			},
		},
	}}
}

func (builder *FuncTemplateBuilder) Param(paramName string, paramDescription string, defaultValues ...interface{}) *FuncTemplateBuilder {
	param := TemplateParam{
		Name:        paramName,
		Description: paramDescription,
	}
	switch len(defaultValues) {
	case 1:
		defaultValueProvider, isProvider := defaultValues[0].(func(ArgMap) interface{})
		if isProvider {
			param.DefaultValueProvider = defaultValueProvider
		} else {
			if reflect.TypeOf(defaultValues[0]).Kind() == reflect.Func {
				panic("default value provider should be func(map[string]interface{})interface{}")
			}
			param.DefaultValueProvider = func(argMap ArgMap) interface{} {
				return defaultValues[0]
			}
		}
	case 0:
		// ignore
	default:
		panic("only one default value should be provided")
	}
	builder.funcTemplate.templateParams[paramName] = param
	return builder
}

func (builder *FuncTemplateBuilder) Generators(kv ...interface{}) *FuncTemplateBuilder {
	for i := 0; i < len(kv); i += 2 {
		k := kv[i].(string)
		v := kv[i+1]
		builder.funcTemplate.generators[k] = v
	}
	return builder
}

func (builder *FuncTemplateBuilder) ImportFunc(funcTemplates ...*FuncTemplate) *FuncTemplateBuilder {
	for _, funcTemplate := range funcTemplates {
		builder.funcTemplate.importedFuncTemplates[funcTemplate.funcName] = funcTemplate
	}
	return builder
}

func (builder *FuncTemplateBuilder) ImportStruct(structTemplates ...*StructTemplate) *FuncTemplateBuilder {
	for _, structTemplate := range structTemplates {
		builder.funcTemplate.importedStructTemplates[structTemplate.structName] = structTemplate
	}
	return builder
}

func (builder *FuncTemplateBuilder) ImportPackage(pkg string) *FuncTemplateBuilder {
	builder.funcTemplate.importedPackages[pkg] = true
	return builder
}

func (builder *FuncTemplateBuilder) Source(source string) *FuncTemplate {
	builder.funcTemplate.templateSource = source
	return builder.funcTemplate
}

type FuncTemplate struct {
	*funcSignature
	definedInFile           string
	templateParams          map[string]TemplateParam
	templateSource          string
	generators              map[string]interface{}
	importedFuncTemplates   map[string]*FuncTemplate
	importedStructTemplates map[string]*StructTemplate
	importedPackages        map[string]bool
}

func (funcTemplate *FuncTemplate) ImportFunc(funcTemplates ...*FuncTemplate) {
	for _, dep := range funcTemplates {
		funcTemplate.importedFuncTemplates[dep.funcName] = dep
	}
}

type TemplateParam struct {
	Name                 string
	Description          string
	DefaultValueProvider func(argMap ArgMap) interface{}
}
