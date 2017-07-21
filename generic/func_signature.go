package generic

import (
	"strings"
	"fmt"
	"bytes"
	"reflect"
)

func parseSignature(input string) (*funcSignature, error) {
	leftBrace := strings.IndexByte(input, '(')
	if leftBrace == -1 {
		return nil, fmt.Errorf("( not found in " + input)
	}
	funcName := strings.TrimSpace(input[:leftBrace])
	rightBrace := strings.IndexByte(input, ')')
	if rightBrace == -1 {
		return nil, fmt.Errorf(") not found in " + input)
	}
	signature := &funcSignature{
		funcName: funcName,
	}
	params := strings.TrimSpace(input[leftBrace+1:rightBrace])
	if len(params) > 0 {
		for _, param := range strings.Split(params, ",") {
			nameAndType := strings.SplitN(strings.TrimSpace(param), " ", 2)
			signature.funcParams = append(signature.funcParams, funcParam{
				paramName: strings.TrimSpace(nameAndType[0]),
				paramType: strings.TrimSpace(nameAndType[1]),
			})
		}
	} else {
		signature.funcParams = []funcParam{}
	}
	returns := strings.TrimFunc(input[rightBrace:], func(r rune) bool {
		switch r {
		case ' ', '\t', '\r', '\n', '(', ')':
			return true
		}
		return false
	})
	if len(returns) > 0 {
		for _, ret := range strings.Split(returns, ",") {
			nameAndType := strings.SplitN(strings.TrimSpace(ret), " ", 2)
			if len(nameAndType) == 1 {
				nameAndType = []string{"", nameAndType[0]}
			}
			signature.funcReturns = append(signature.funcReturns, funcReturn{
				returnName: strings.TrimSpace(nameAndType[0]),
				returnType: strings.TrimSpace(nameAndType[1]),
			})
		}
	} else {
		signature.funcReturns = []funcReturn{}
	}
	return signature, nil
}

type funcSignature struct {
	funcName    string
	funcParams  []funcParam
	funcReturns []funcReturn
}

type funcParam struct {
	paramName string
	paramType string
}

type funcReturn struct {
	returnName string
	returnType string
}

func (signature *funcSignature) expand(out *bytes.Buffer,
	expandedFuncName string, argMap map[string]interface{}) {
	out.WriteString("\nfunc ")
	out.WriteString(expandedFuncName)
	out.WriteByte('(')
	for i, param := range signature.funcParams {
		if i != 0 {
			out.WriteByte(',')
		}
		out.WriteString(param.paramName)
		out.WriteByte(' ')
		typ, isType := argMap[param.paramType].(reflect.Type)
		if isType {
			if state.testMode {
				out.WriteString(internalizeType(typ))
			} else {
				out.WriteString(genName(typ))
			}
		} else {
			out.WriteString(param.paramType)
		}
	}
	out.WriteByte(')')
	out.WriteByte('(')
	for i, ret := range signature.funcReturns {
		if i != 0 {
			out.WriteByte(',')
		}
		out.WriteString(ret.returnName)
		out.WriteByte(' ')
		typ, isType := argMap[ret.returnType].(reflect.Type)
		if isType {
			if state.testMode {
				out.WriteString(internalizeType(typ))
			} else {
				out.WriteString(genName(typ))
			}
		} else {
			out.WriteString(ret.returnType)
		}
	}
	out.WriteByte(')')
	out.WriteByte('{')
}

func (funcTemplate *FuncTemplate) expandTestModeEntryFunc(out *bytes.Buffer,
	expandedFuncName string, argMap map[string]interface{}) error {
	signature := funcTemplate.funcSignature
	invocation := bytes.NewBuffer(nil)

	typedFuncName, err := funcTemplate.expandTestModeInternalFunc(argMap)
	if err != nil {
		return err
	}
	if len(funcTemplate.funcReturns) > 0 {
		invocation.WriteString("return ")
	}
	invocation.WriteString(typedFuncName)
	invocation.WriteByte('(')

	out.WriteString("\nfunc ")
	out.WriteString(expandedFuncName)
	out.WriteByte('(')
	for i, param := range signature.funcParams {
		if i != 0 {
			out.WriteByte(',')
			invocation.WriteByte(',')
		}
		out.WriteString(param.paramName)
		out.WriteByte(' ')
		typ, isType := argMap[param.paramType].(reflect.Type)
		if isType {
			out.WriteString("interface{}")
			typeName, err := genCast(param.paramName, typ)
			if err != nil {
				return err
			}
			invocation.WriteString(typeName)
		} else {
			out.WriteString(param.paramType)
			invocation.WriteString(param.paramName)
		}
	}
	out.WriteByte(')')
	out.WriteByte('(')
	for i, ret := range signature.funcReturns {
		if i != 0 {
			out.WriteByte(',')
		}
		out.WriteString(ret.returnName)
		out.WriteByte(' ')
		_, isType := argMap[ret.returnType].(reflect.Type)
		if isType {
			out.WriteString("interface{}")
		} else {
			out.WriteString(ret.returnType)
		}
	}
	out.WriteString(") {\n")
	invocation.WriteByte(')')
	out.Write(invocation.Bytes())
	out.WriteString("\n}")
	return nil
}

func (funcTemplate *FuncTemplate) expandTestModeInternalFunc(argMap map[string]interface{}) (string, error) {
	templateArgsWithoutAsEmptyInterface := []interface{}{}
	for k, v := range argMap {
		if k != "testMode" {
			templateArgsWithoutAsEmptyInterface = append(templateArgsWithoutAsEmptyInterface, k)
			templateArgsWithoutAsEmptyInterface = append(templateArgsWithoutAsEmptyInterface, v)
		}
	}
	state.testMode = true
	return funcTemplate.expand(templateArgsWithoutAsEmptyInterface)
}

func genCast(identifier string, typ reflect.Type) (string, error) {
	state.importPackages["unsafe"] = true
	state.declarations[`
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
func objPtr(obj interface{}) unsafe.Pointer {
	return (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
}
	`] = true
	return fmt.Sprintf("*(*%s)(objPtr(%s))", internalizeType(typ), identifier), nil
}