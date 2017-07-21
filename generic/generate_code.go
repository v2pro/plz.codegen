package generic

import (
	"bytes"
	"io/ioutil"
	"fmt"
)

func GenerateCode(gopath string, pkgPath string) {
	state.out = bytes.NewBuffer(nil)
	state.importPackages = map[string]bool{}
	state.declarations = map[string]bool{}
	state.expandedFuncNames = map[string]bool{}
	state.pkgPath = pkgPath
	for _, funcDeclaration := range funcDeclarations {
		_, err := funcDeclaration.funcTemplate.expand(funcDeclaration.templateArgs)
		if err != nil {
			panic(err.Error())
		}
	}
	prelog := `
package model
	`
	for importPackage := range state.importPackages {
		prelog = fmt.Sprintf(`
%s
import "%s"`, prelog, importPackage)
	}
	for declaration := range state.declarations {
		prelog = prelog + "\n" + declaration
	}
	err := ioutil.WriteFile(gopath + "/src/" + pkgPath + "/generated.go", append([]byte(prelog), state.out.Bytes()...), 0666)
	if err != nil {
		panic(err.Error())
	}
}