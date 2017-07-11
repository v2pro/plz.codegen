package fp_compare

import (
	"plugin"
	"os"
	"io/ioutil"
	"os/exec"
	"bytes"
	"sync"
	"github.com/v2pro/plz"
	"text/template"
	"reflect"
)

var logger = plz.LoggerOf("package", "fp_compare")

type funcTemplate struct {
	variables    map[string]string
	source       string
	funcName     string
	dependencies []*funcTemplate
}

type generator struct {
	generatedTypes map[reflect.Type]bool
}

func (g *generator) gen(fTmpl *funcTemplate, kv ...interface{}) (string, string) {
	generatedSource := ""
	for _, dep := range fTmpl.dependencies {
		_, depSource := g.gen(dep, kv...)
		generatedSource += depSource
	}
	data := map[string]interface{}{}
	for i := 0; i < len(kv); i += 2 {
		data[kv[i].(string)] = kv[i+1]
		typ, _ := kv[i+1].(reflect.Type)
		if typ != nil && typ.Kind() == reflect.Struct {
			if !g.generatedTypes[typ] {
				g.generatedTypes[typ] = true
				generatedSource += generateStruct(typ)
			}
		}
	}
	funcName := genFuncName(fTmpl.funcName, data)
	data["funcName"] = funcName
	tmpl, err := template.New("source").Funcs(map[string]interface{}{
		"name": func_name,
		"cast": func_cast,
	}).Parse(fTmpl.source)
	panicOnError(err)
	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	panicOnError(err)
	return funcName, generatedSource + out.String()
}

func gen(fTmpl *funcTemplate, kv ...interface{}) (string, string) {
	return (&generator{
		generatedTypes: map[reflect.Type]bool{},
	}).gen(fTmpl, kv...)
}

func genFuncName(funcNameTmpl string, data interface{}) string {
	tmpl, err := template.New("funcName").Funcs(map[string]interface{}{
		"name": func_name,
	}).Parse(funcNameTmpl)
	panicOnError(err)
	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	panicOnError(err)
	return out.String()
}

func genSource(template *funcTemplate, kv ...interface{}) string {
	_, src := gen(template, kv...)
	return src
}

var compilerMutex = &sync.Mutex{}

func compile(source string, funcName string) plugin.Symbol {
	compilerMutex.Lock()
	defer compilerMutex.Unlock()
	source = `
package main
import "unsafe"
import "reflect"
	` + source + genObjPtr
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
		logger.Error("compile failed", "source", source)
		panic("failed to compile generated plugin: " + err.Error() + ", " + errBuf.String())
	}
	generatedPlugin, err := plugin.Open(soFileName)
	if err != nil {
		panic("failed to load generated plugin: " + err.Error())
	}
	compareObj, err := generatedPlugin.Lookup(funcName)
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
	return compareObj
}
