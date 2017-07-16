package gen

import (
	"bytes"
	"crypto/sha1"
	"encoding/base32"
	"io/ioutil"
	"os"
	"os/exec"
	"plugin"
	"reflect"
	"runtime"
	"sync"
	"fmt"
)

var compilePluginMutex = &sync.Mutex{}

// CompilePlugin compiles the expanded template into plugin, producing a .so file
func CompilePlugin(soFileName string) {
	compilePluginMutex.Lock()
	defer compilePluginMutex.Unlock()
	generator := newGenerator()
	source := ""
	for _, expansion := range expansions {
		_, oneSource := generator.gen(expansion.template, expansion.templateArgs...)
		source += oneSource
	}
	logger.Debug("generated source", "source", source)
	compileAndOpenPlugin(soFileName, source)
}
func newGenerator() *generator {
	return &generator{
		generatedTypes:        map[reflect.Type]bool{},
		generatedFuncs:        map[string]bool{},
		generatedDeclarations: map[string]bool{},
	}
}

func compileAndOpenPlugin(origSoFileName string, origSource string) *plugin.Plugin {
	prelog := `
package main
import "unsafe"
import "fmt"
import "io"
`
	for pkg := range ImportPackages {
		prelog += "import \"" + pkg + "\"\n"
	}
	sourceHash := hash(origSource)
	prelog += fmt.Sprintf(`
var SOURCE__HASH = "%s"
var ioEOF = io.EOF
var debugLog = fmt.Println

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
	`, sourceHash)
	source := prelog + origSource
	fileName := sourceHash
	homeDir := os.Getenv("HOME")
	cacheDir := homeDir + "/.wombat/"
	//cacheDir = "/tmp/"
	if _, err := os.Stat(cacheDir); err != nil {
		os.Mkdir(cacheDir, 0777)
	}
	srcFileName := cacheDir + fileName + ".go"
	soFileName := origSoFileName
	if soFileName == "" {
		soFileName = cacheDir + fileName + ".so"
	}
	if _, err := os.Stat(soFileName); err != nil {
		err := ioutil.WriteFile(srcFileName, []byte(source), 0666)
		if err != nil {
			panic("failed to generate source code: " + err.Error())
		}
		logger.Debug("build plugin", "soFileName", soFileName, "srcFileName", srcFileName)
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soFileName, srcFileName)
		var errBuf bytes.Buffer
		cmd.Stderr = &errBuf
		var outBuf bytes.Buffer
		cmd.Stdout = &outBuf
		err = cmd.Run()
		if err != nil {
			logger.Error("compile failed", "srcFileName", srcFileName, "source", annotateLines(source))
			panic("failed to compile generated plugin: " + err.Error() + ", " + errBuf.String())
		}
	}
	logger.Debug("open plugin", "soFileName", soFileName)
	thePlugin, err := plugin.Open(soFileName)
	if err != nil {
		panic("failed to load generated plugin: " + err.Error())
	}
	if origSoFileName != "" && !verifySourceHash(thePlugin, sourceHash) {
		logger.Info("remove out")
		err = os.Remove(origSoFileName)
		if err != nil {
			panic("failed to removed out of date plugin: " + err.Error())
		}
		return compileAndOpenPlugin(origSoFileName, origSource)
	}
	return thePlugin
}

func verifySourceHash(thePlugin *plugin.Plugin, sourceHash string) bool {
	symbol, err := thePlugin.Lookup("SOURCE__HASH")
	if err != nil {
		logger.Error("SOURCE__HASH missing from so")
		return false
	}
	actualSourceHash, isStr := symbol.(*string)
	if !isStr {
		logger.Error("SOURCE__HASH is not string")
		return false
	}
	if *actualSourceHash != sourceHash {
		logger.Error("SOURCE__HASH mismatch", "expected", sourceHash, "actual", *actualSourceHash)
		return false
	}
	return true
}

func hash(source string) string {
	h := sha1.New()
	h.Write([]byte(source))
	h.Write([]byte(runtime.Version()))
	return "g" + base32.StdEncoding.EncodeToString(h.Sum(nil))
}
