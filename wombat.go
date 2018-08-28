package wombat

import (
	"bytes"
	"reflect"
	"os/exec"
	"fmt"
	"io/ioutil"
	"os"
)

var Int = reflect.TypeOf(int(0))
var Int8 = reflect.TypeOf(int8(0))
var Int16 = reflect.TypeOf(int16(0))
var Int32 = reflect.TypeOf(int32(0))
var Int64 = reflect.TypeOf(int64(0))
var Uint = reflect.TypeOf(uint(0))
var Uint8 = reflect.TypeOf(uint8(0))
var Uint16 = reflect.TypeOf(uint16(0))
var Uint32 = reflect.TypeOf(uint32(0))
var Uint64 = reflect.TypeOf(uint64(0))
var Float32 = reflect.TypeOf(float32(0))
var Float64 = reflect.TypeOf(float64(0))
var String = reflect.TypeOf("")
var Bool = reflect.TypeOf(true)

func Codegen(pkgPath string) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		reportError("GOPATH env not found")
	}
	writeCodeGeneratorMain(gopath, pkgPath)
	executeTmpCodegen(gopath + "/bin/tmp-codegen")
}

func writeCodeGeneratorMain(gopath string, pkgPath string) {
	dir := gopath+"/src/tmp-codegen"
	if _, err := os.Stat(dir); err != nil {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			reportError(err.Error())
		}
	}
	ioutil.WriteFile(dir+"/main.go", []byte(fmt.Sprintf(`
package main
import _ "%s"
import "github.com/v2pro/wombat/generic"
func main() {
	generic.GenerateCode("%s", "%s")
}
	`, pkgPath, gopath, pkgPath)), 0666)
	goInstallTmpCodegen()
}

func goInstallTmpCodegen() {
	cmd := exec.Command("go", "install", "-tags", "codegen", "tmp-codegen")
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	err := cmd.Run()
	if err != nil {
		fmt.Println(errBuf.String())
		fmt.Println(outBuf.String())
		reportError(err.Error())
	}
}

func executeTmpCodegen(file string) {
	cmd := exec.Command(file)
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	err := cmd.Run()
	if err != nil {
		fmt.Println(errBuf.String())
		fmt.Println(outBuf.String())
		reportError(err.Error())
	}
	fmt.Println(errBuf.String())
	fmt.Println(outBuf.String())
}

func reportError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

