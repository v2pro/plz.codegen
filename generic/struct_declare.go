package generic

import "reflect"

func DeclareStruct(structTemplate *StructTemplate, intefaceType reflect.Type) {
	ctorTemplate := structTemplate.ctorTemplate()
	DeclareFunc(ctorTemplate, "I", intefaceType)
}