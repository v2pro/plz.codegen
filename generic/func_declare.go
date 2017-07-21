package generic

type funcDeclaration struct {
	funcTemplate *FuncTemplate
	templateArgs []interface{}
}

var funcDeclarations = []funcDeclaration{}

func DeclareFunc(funcTemplate *FuncTemplate, templateArgs ...interface{}) {
	declaration := funcDeclaration{
		funcTemplate: funcTemplate,
		templateArgs: templateArgs,
	}
	funcDeclarations = append(funcDeclarations, declaration)
}
