package gen

import (
	"io/ioutil"
	"fmt"
)

var registeredFuncs = map[string]interface{}{}

func GenerateCode(path string) {
	generator := newGenerator()
	source := ""
	funcNames := []string{}
	for _, expansion := range declaredExpansions {
		funcName, oneSource := generator.gen(expansion.template, expansion.templateArgs...)
		source += oneSource
		funcNames = append(funcNames, funcName)
	}
	prelog := `
package model
import "unsafe"
import "fmt"
import "io"
import "github.com/v2pro/wombat/gen"
`
	for pkg := range ImportPackages {
		prelog += "import \"" + pkg + "\"\n"
	}
	prelog += `
var ioEOF = io.EOF
var debugLog = fmt.Println

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
func init() {
	`
	for _, funcName := range funcNames {
		prelog += fmt.Sprintf(`gen.RegisterFunc("Exported_%s", Exported_%s)`, funcName, funcName)
		prelog += "\n"
	}
	prelog += `}`
	err := ioutil.WriteFile(path, []byte(prelog+source), 0666)
	if err != nil {
		panic(err.Error())
	}
}

func lookupFunc(funcName string) interface{} {
	f := registeredFuncs[funcName]
	if f != nil {
		return f
	}
	return lookupFuncFromPlugins(funcName)
}

func RegisterFunc(funcName string, funcObj interface{}) {
	registeredFuncs[funcName] = funcObj
}
