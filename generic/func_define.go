package generic

type FuncTemplateBuilder struct {
	funcTemplate *FuncTemplate
}

func Func(funcName string) *FuncTemplateBuilder {
	return &FuncTemplateBuilder{funcTemplate: &FuncTemplate{
		funcName: funcName,
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

func (builder *FuncTemplateBuilder) Source(source string) *FuncTemplate {
	builder.funcTemplate.templateSource = source
	return builder.funcTemplate
}

type FuncTemplate struct {
	funcName string
	templateParams map[string]interface{}
	templateSource string
}