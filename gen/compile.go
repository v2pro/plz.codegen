package gen

import (
	"bytes"
	"github.com/v2pro/plz"
	"github.com/v2pro/plz/logging"
	"plugin"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"fmt"
)

var logger = plz.LoggerOf("package", "gen")

func init() {
	logging.Providers = append(logging.Providers, func(loggerKv []interface{}) logging.Logger {
		for i := 0; i < len(loggerKv); i += 2 {
			key := loggerKv[i].(string)
			if key == "package" && "gen" == loggerKv[i+1] {
				return logging.NewStderrLogger(loggerKv, logging.LEVEL_DEBUG)
			}
		}
		return nil
	})
}

// FuncTemplate used to generate similar functions with template by applying different arguments
type FuncTemplate struct {
	FuncTemplateName string
	Dependencies   []*FuncTemplate
	TemplateParams map[string]string
	Source         string
	FuncName       string
	GenMap         map[string]interface{}
	Declarations   string
}

func (template *FuncTemplate) AddDependency(dep *FuncTemplate) {
	template.Dependencies = append(template.Dependencies, dep)
}

type generator struct {
	generatedTypes        map[reflect.Type]bool
	generatedFuncs        map[string]bool
	generatedDeclarations map[string]bool
}

func (g *generator) gen(fTmpl *FuncTemplate, args ...interface{}) (string, string) {
	generatedSource := ""
	if fTmpl.Declarations != "" && !g.generatedDeclarations[fTmpl.Declarations] {
		generatedSource += fTmpl.Declarations
		g.generatedDeclarations[fTmpl.Declarations] = true
	}
	templateArgs := map[string]interface{}{}
	templateParams := map[string]string{}
	for k, v := range fTmpl.TemplateParams {
		templateParams[k] = v
	}
	for i := 0; i < len(args); i += 2 {
		param := args[i].(string)
		_, isDeclared := templateParams[param]
		if !isDeclared {
			logger.Error("variable not declared", "param", param, "args", args)
			panic("variable " + param + " is not declared")
		}
		delete(templateParams, param)
		typ, _ := args[i+1].(reflect.Type)
		if typ != nil {
			args[i+1] = typ
			generatedSource += g.genTypeDef(typ)
		}
		templateArgs[param] = args[i+1]
	}
	for k, v := range templateParams {
		logger.Error("missing variable", "varName", k, "varDescription", v, "args", args)
		panic("missing variable " + k + ": " + v)
	}
	funcName := genFuncName(fTmpl.FuncName, templateArgs)
	if g.generatedFuncs[funcName] {
		return funcName, ""
	}
	templateArgs["funcName"] = funcName
	depMap := map[string]*FuncTemplate{}
	for _, dep := range fTmpl.Dependencies {
		depMap[dep.FuncTemplateName] = dep
	}
	genMap := map[string]interface{}{
		"gen": func(depName string, newKv ...interface{}) string {
			dep := depMap[depName]
			if dep == nil {
				logger.Error("referenced unfound dependency", "depName", depName, "kv", newKv)
				panic("referenced unfound dependency " + depName)
			}
			funcName, source := g.gen(dep, newKv...)
			generatedSource += source
			return funcName
		},
		"cast": func(identifier string, typ reflect.Type) string {
			if typ.Kind() == reflect.Interface {
				if typ.NumMethod() == 0 {
					return identifier
				}
				return fmt.Sprintf("%s.(%s)", identifier, funcGetName(typ))
			}
			objPtrFuncName, objPtrSource := g.gen(objPtrF, "T", typ)
			generatedSource += objPtrSource
			if typ.Kind() == reflect.Ptr {
				return "((" + funcGetName(typ) + ")(" + objPtrFuncName + "(" + identifier + ")))"
			}
			return "(*(*" + funcGetName(typ) + ")(" + objPtrFuncName + "(" + identifier + ")))"
		},
	}
	fillDefaultFuncMap(genMap)
	for k, v := range fTmpl.GenMap {
		genMap[k] = v
	}
	out := executeTemplate(fTmpl.Source, genMap, templateArgs)
	g.generatedFuncs[funcName] = true
	return funcName, generatedSource + out + "\n// generated from " + fTmpl.FuncTemplateName + "\n"
}

func executeTemplate(tmplSource string, funcMap map[string]interface{}, data map[string]interface{}) string {
	defer func() {
		recovered := recover()
		if recovered != nil {
			logger.Error("generate source failed", "tmplSource", annotateLines(tmplSource))
			panic(recovered)
		}
	}()
	tmpl, err := template.New(hash(tmplSource)).Funcs(funcMap).Parse(tmplSource)
	panicOnError(err)
	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	panicOnError(err)
	return out.String()
}

func fillDefaultFuncMap(funcMap map[string]interface{}) {
	funcMap["fieldOf"] = funcFieldOf
	funcMap["elem"] = funcElem
	funcMap["isPtr"] = funcIsPtr
	funcMap["name"] = funcGetName
	funcMap["symbol"] = funcSymbol
}

func gen(fTmpl *FuncTemplate, kv ...interface{}) (string, string) {
	return newGenerator().gen(fTmpl, kv...)
}

func genFuncName(funcNameTmpl string, data interface{}) string {
	tmpl, err := template.New(hash(funcNameTmpl)).Funcs(map[string]interface{}{
		"symbol": funcSymbol,
		"name":   funcGetName,
		"elem":   funcElem,
	}).Parse(funcNameTmpl)
	panicOnError(err)
	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	panicOnError(err)
	return out.String()
}

type expansion struct {
	template *FuncTemplate
	templateArgs       []interface{}
}

var declaredExpansions = []expansion{}

func Declare(template *FuncTemplate, templateArgs ...interface{}) {
	declaredExpansions = append(declaredExpansions,
		expansion{template: template, templateArgs: templateArgs})
}

// Expand expand the function template with provided type arguments,
// compiles the code and loads as executable
func Expand(template *FuncTemplate, templateArgs ...interface{}) plugin.Symbol {
	funcName, source := gen(template, templateArgs...)
	logger.Debug("generated source", "source", source)
	symbol := lookupFunc("Exported_" + funcName)
	if symbol != nil {
		return symbol
	}
	assertDynamicCompilation(template, templateArgs)
	return dynamicCompile("Exported_"+funcName, source)
}

var dynamicCompileMutex = &sync.Mutex{}

func dynamicCompile(funcName, source string) plugin.Symbol {
	dynamicCompileMutex.Lock()
	defer dynamicCompileMutex.Unlock()
	thePlugin := compileAndOpenPlugin("", source)
	symbol, err := thePlugin.Lookup(funcName)
	if err != nil {
		panic("failed to lookup symbol from generated plugin: " + err.Error())
	}
	return symbol
}

func annotateLines(source string) string {
	var buf bytes.Buffer
	lines := strings.Split(source, "\n")
	for i, line := range lines {
		lineNo := strconv.FormatInt(int64(i+1), 10)
		buf.WriteString(lineNo)
		buf.WriteString(": ")
		buf.WriteString(line)
		buf.WriteString("\n")
	}
	return buf.String()
}
