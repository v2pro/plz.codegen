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

// FuncTemplate defines a generic function using template
type FuncTemplate struct {
	Variables    map[string]string
	Source       string
	FuncName     string
	Dependencies map[string]*FuncTemplate
	FuncMap      map[string]interface{}
}

type generator struct {
	generatedTypes map[reflect.Type]bool
	generatedFuncs map[string]bool
}

func (g *generator) gen(fTmpl *FuncTemplate, args ...interface{}) (string, string) {
	generatedSource := ""
	data := map[string]interface{}{}
	variables := map[string]string{}
	for k, v := range fTmpl.Variables {
		variables[k] = v
	}
	for i := 0; i < len(args); i += 2 {
		varName := args[i].(string)
		_, isDeclared := variables[varName]
		if !isDeclared {
			logger.Error("variable not declared", "varName", varName, "args", args)
			panic("variable " + varName + " is not declared")
		}
		delete(variables, varName)
		data[varName] = args[i+1]
		typ, _ := args[i+1].(reflect.Type)
		if typ != nil && (typ.Kind() == reflect.Struct || typ.Kind() == reflect.Ptr) {
			generatedSource += g.genStruct(typ)
		}
	}
	for k, v := range variables {
		logger.Error("missing variable", "varName", k, "varDescription", v, "args", args)
		panic("missing variable " + k + ": " + v)
	}
	funcName := genFuncName(fTmpl.FuncName, data)
	if g.generatedFuncs[funcName] {
		return funcName, ""
	}
	data["funcName"] = funcName
	funcMap := map[string]interface{}{
		"gen": func(depName string, newKv ...interface{}) interface{} {
			dep := fTmpl.Dependencies[depName]
			if dep == nil {
				panic("referenced unfound dependency " + depName)
			}
			funcName, source := g.gen(dep, newKv...)
			return struct {
				FuncName string
				Source   string
			}{FuncName: funcName, Source: source}
		},
		"cast": func(identifier string, typ reflect.Type) string {
			objPtrFuncName, objPtrSource := g.gen(objPtrF, "T", typ)
			generatedSource += objPtrSource
			if typ.Kind() == reflect.Ptr {
				return "((" + funcGetName(typ) + ")(" + objPtrFuncName + "(" + identifier + ")))"
			}
			return "(*(*" + funcGetName(typ) + ")(" + objPtrFuncName + "(" + identifier + ")))"

		},
	}
	fillDefaultFuncMap(funcMap)
	for k, v := range fTmpl.FuncMap {
		funcMap[k] = v
	}
	out := executeTemplate(fTmpl.Source, funcMap, data)
	g.generatedFuncs[funcName] = true
	return funcName, generatedSource + out
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
	return (&generator{
		generatedTypes: map[reflect.Type]bool{},
		generatedFuncs: map[string]bool{},
	}).gen(fTmpl, kv...)
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

type compileOp struct {
	template *FuncTemplate
	kv       []interface{}
}

// Compile expand the function template with provided type arguments,
// compiles the code and loads as executable
func Compile(template *FuncTemplate, kv ...interface{}) plugin.Symbol {
	if isInBatchCompileMode {
		panic(&compileOp{template: template, kv: kv})
	}
	funcName, source := gen(template, kv...)
	logger.Debug("generated source", "source", source)
	symbol := lookupFunc("Exported_" + funcName)
	if symbol != nil {
		return symbol
	}
	return dynamicCompile("Exported_"+funcName, source)
}

var dynamicCompileMutex = &sync.Mutex{}

func dynamicCompile(funcName, source string) plugin.Symbol {
	if dynamicCompilationDisabled {
		logger.Error("dynamic compilation disabled", "funcName", funcName, "source", source)
		panic("dynamic compilation disabled")
	}
	dynamicCompileMutex.Lock()
	defer dynamicCompileMutex.Unlock()
	thePlugin := compileAndOpenPlugin(source)
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

var dynamicCompilationDisabled = false

// DisableDynamicCompilation prevents dynamic compilation, everything should be loaded from LoadPlugin
func DisableDynamicCompilation() {
	dynamicCompilationDisabled = true
}
