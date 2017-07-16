package gen

import (
	"reflect"
	"fmt"
	"strings"
)

var dynamicCompilationDisabled = false

// DisableDynamicCompilation prevents dynamic compilation, everything should be loaded from LoadPlugin
func DisableDynamicCompilation() {
	dynamicCompilationDisabled = true
}

func assertDynamicCompilation(template *FuncTemplate, templateArgs []interface{}) {
	if !dynamicCompilationDisabled {
		return
	}
	logger.Error("dynamic compilation disabled",
		"templateFuncName", template.FuncTemplateName,
		"templateArgs", templateArgs)
	argsAsStr := []string{}
	for _, arg := range templateArgs {
		switch typedArg := arg.(type) {
		case string:
			argsAsStr = append(argsAsStr, `"`+typedArg+`"`)
		case reflect.Type:
			argsAsStr = append(argsAsStr, reverseTypeToStr(typedArg))
		default:
			argsAsStr = append(argsAsStr, fmt.Sprintf("%v", argsAsStr))
		}
	}
	panic(fmt.Sprintf("please add wombat.Expand(%s.F, %s) to init() or enable dynamic compilation",
		template.FuncTemplateName,
		strings.Join(argsAsStr, ", ")))
}

func reverseTypeToStr(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int:
		return "wombat.Int"
	case reflect.Int8:
		return "wombat.Int8"
	case reflect.Int16:
		return "wombat.Int16"
	case reflect.Int32:
		return "wombat.Int32"
	case reflect.Int64:
		return "wombat.Int64"
	case reflect.Uint:
		return "wombat.Uint"
	case reflect.Uint8:
		return "wombat.Uint8"
	case reflect.Uint16:
		return "wombat.Uint16"
	case reflect.Uint32:
		return "wombat.Uint32"
	case reflect.Uint64:
		return "wombat.Uint64"
	case reflect.Float32:
		return "wombat.Float32"
	case reflect.Float64:
		return "wombat.Float64"
	case reflect.String:
		return "wombat.String"
	case reflect.Bool:
		return "wombat.Bool"
	case reflect.Struct:
		return "reflect.TypeOf(" + typ.String() + "{})"
	}
	panic("not implemented")
}