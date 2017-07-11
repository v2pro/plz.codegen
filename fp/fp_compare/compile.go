package fp_compare

import (
	"plugin"
	"os"
	"io/ioutil"
	"os/exec"
	"strings"
	"bytes"
	"sync"
	"github.com/v2pro/plz"
)

var logger = plz.LoggerOf("package", "fp_compare")

type funcTemplate struct {
	variables map[string]string
	source    string
	funcName  string
	dependencies []*funcTemplate
}

func render(template *funcTemplate, kv ...string) (funcName string, src string) {
	funcName = template.funcName
	src = template.source
	// TODO: check all variables defined
	for i := 0; i < len(kv); i += 2 {
		k := kv[i]
		v := kv[i+1]
		kNoDot := k + "|nodot"
		vNoDot := strings.Replace(v, ".", "__", -1)
		src = strings.Replace(src, "{{"+k+"}}", v, -1)
		funcName = strings.Replace(funcName, "{{"+k+"}}", v, -1)
		src = strings.Replace(src, "{{"+kNoDot+"}}", vNoDot, -1)
		funcName = strings.Replace(funcName, "{{"+kNoDot+"}}", vNoDot, -1)
	}
	src = strings.Replace(src, "{{funcName}}", funcName, -1)
	return
}

func renderSource(template *funcTemplate, kv ...string) string {
	_, src := render(template, kv...)
	return src
}

var compilerMutex = &sync.Mutex{}

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
