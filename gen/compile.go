package gen

import (
	"bytes"
	"github.com/v2pro/plz"
	"github.com/v2pro/plz/logging"
	"io/ioutil"
	"os"
	"os/exec"
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
			if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Map {
				return "((" + funcGetName(typ) + ")(" + objPtrFuncName + "(" + identifier + ")))"
			} else {
				return "(*(*" + funcGetName(typ) + ")(" + objPtrFuncName + "(" + identifier + ")))"
			}
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
	tmpl, err := template.New(NewID().String()).Funcs(funcMap).Parse(tmplSource)
	panicOnError(err)
	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	panicOnError(err)
	return out.String()
}

func fillDefaultFuncMap(funcMap map[string]interface{}) {
	funcMap["isOnePtrStructOrArray"] = funcIsOnePtrStructOrArray
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
	tmpl, err := template.New(NewID().String()).Funcs(map[string]interface{}{
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

func Compile(template *FuncTemplate, kv ...interface{}) plugin.Symbol {
	if isInBatchCompileMode {
		panic(&compileOp{template: template, kv: kv})
	}
	funcName, source := gen(template, kv...)
	logger.Debug("generated source", "source", source)
	symbol := lookupFunc(funcName)
	if symbol != nil {
		return symbol
	}
	return dynamicCompile(funcName, source)
}

const prelog = `
package main
import "unsafe"

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
`

var dynamicCompileMutex = &sync.Mutex{}
func dynamicCompile(funcName, source string) plugin.Symbol {
	if dynamicCompilationDisabled {
		logger.Error("dynamic compilation disabled", "funcName", funcName, "source", source)
		panic("dynamic compilation disabled")
	}
	dynamicCompileMutex.Lock()
	defer dynamicCompileMutex.Unlock()
	source = prelog + source
	srcFileName := "/tmp/" + NewID().String() + ".go"
	soFileName := "/tmp/" + NewID().String() + ".so"
	err := ioutil.WriteFile(srcFileName, []byte(source), 0666)
	if err != nil {
		panic("failed to generate source code: " + err.Error())
	}
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soFileName, srcFileName)
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	err = cmd.Run()
	if err != nil {
		logger.Error("compile failed", "source", annotateLines(source))
		panic("failed to compile generated plugin: " + err.Error() + ", " + errBuf.String())
	}
	generatedPlugin, err := plugin.Open(soFileName)
	if err != nil {
		panic("failed to load generated plugin: " + err.Error())
	}
	symbol, err := generatedPlugin.Lookup(funcName)
	if err != nil {
		panic("failed to lookup symbol from generated plugin: " + err.Error())
	}
	err = os.Remove(srcFileName)
	if err != nil {
		logger.Error("failed to remove generated source", "srcFileName", srcFileName)
	}
	err = os.Remove(soFileName)
	if err != nil {
		logger.Error("failed to remove generated plugin", "soFileName", soFileName)
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

func DisableDynamicCompilation() {
	dynamicCompilationDisabled = true
}
