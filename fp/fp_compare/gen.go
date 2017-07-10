package fp_compare

import (
	"io/ioutil"
	"os/exec"
	"bytes"
	"plugin"
	"strings"
	"reflect"
	"os"
	"github.com/v2pro/plz"
	"sync"
)

var logger = plz.LoggerOf("package", "fp_compare")
var compilerMutex = &sync.Mutex{}

type funcTemplate struct {
	variables map[string]string
	source    string
	funcName  string
}

var compareSymbols = struct {
	template funcTemplate
	cache map[reflect.Type]func(interface{}, interface{}) int
}{
	cache: map[reflect.Type]func(interface{}, interface{}) int{},
	template: funcTemplate{
		variables: map[string]string{
			"T": "the type to compare",
		},
		source: `
func Compare_{{T}}(obj1 interface{}, obj2 interface{}) int {
	return typed_Compare_{{T}}(obj1.({{T}}), obj2.({{T}}))
}
func typed_Compare_{{T}}(obj1 {{T}}, obj2 {{T}}) int {
	if (obj1 < obj2) {
		return -1
	} else if (obj1 == obj2) {
		return 0
	} else {
		return 1
	}
}`,
		funcName: `Compare_{{T}}`,
	},
}

func render(template string, kv ...string) string {
	for i := 0; i < len(kv); i+=2 {
		template = strings.Replace(template, "{{" + kv[i] + "}}", kv[i+1], -1)
	}
	return template
}

func Compare(obj1 interface{}, obj2 interface{}) int {
	typ := reflect.TypeOf(obj1)
	compare := compareSymbols.cache[typ]
	if compare == nil {
		typeName := typ.String()
		source := render(compareSymbols.template.source, `T`, typeName)
		funcName := render(compareSymbols.template.funcName, `T`, typeName)
		compareObj := compile(source, funcName)
		compare = compareObj.(func(interface{}, interface{}) int)
		compareSymbols.cache[typ] = compare
	}
	return compare(obj1, obj2)
}

func compile(source string, funcName string) plugin.Symbol {
	compilerMutex.Lock()
	defer compilerMutex.Unlock()
	source = "package main\n" + source
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
