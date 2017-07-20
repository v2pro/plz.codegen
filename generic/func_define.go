package generic

import (
	"fmt"
)

type FuncTemplateBuilder struct {
	funcTemplate *FuncTemplate
}

func Func(funcName string) *FuncTemplateBuilder {
	importedFuncTemplates := map[string]*FuncTemplate{}
	return &FuncTemplateBuilder{funcTemplate: &FuncTemplate{
		funcName: funcName,
		templateParams: map[string]interface{}{},
		importedFuncTemplates: importedFuncTemplates,
		generators: map[string]interface{}{
			"name": genName,
			"expand": func(depName string, templateArgs ...interface{}) (string, error) {
				dep := importedFuncTemplates[depName]
				if dep == nil {
					logger.Error(nil, "missing dependency", "depName", depName)
					return "", fmt.Errorf(
						"referenced generic function %s should be imported by ImportFunc",
						depName)
				}
				expandedFuncName, _, err := dep.expand(templateArgs)
				if err != nil {
					return "", logger.Error(err, fmt.Sprintf("expand %s failed", depName),
					"templateArgs", templateArgs)
				}
				return expandedFuncName, nil
			},
		},
	}}
}

func (builder *FuncTemplateBuilder) Params(kv ...interface{}) *FuncTemplateBuilder {
	for i := 0; i < len(kv); i+=2 {
		k := kv[i].(string)
		v := kv[i+1]
		builder.funcTemplate.templateParams[k] = v
	}
	return builder
}

func (builder *FuncTemplateBuilder) Generators(kv ...interface{}) *FuncTemplateBuilder {
	for i := 0; i < len(kv); i+=2 {
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

func (builder *FuncTemplateBuilder) Source(source string) *FuncTemplate {
	builder.funcTemplate.templateSource = source
	return builder.funcTemplate
}

type FuncTemplate struct {
	funcName string
	templateParams map[string]interface{}
	templateSource string
	generators map[string]interface{}
	importedFuncTemplates map[string]*FuncTemplate
}