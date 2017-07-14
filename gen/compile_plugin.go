package gen

import (
	"sync"
	"fmt"
	"os/exec"
	"io/ioutil"
	"reflect"
	"bytes"
	"os"
)

var compilePluginMutex = &sync.Mutex{}
var isInBatchCompileMode = false

func CompilePlugin(soFileName string, compileOpTriggers ...func()) {
	compilePluginMutex.Lock()
	defer compilePluginMutex.Unlock()
	isInBatchCompileMode = true
	defer func() {
		isInBatchCompileMode = false
	}()
	compileOps := []*compileOp{}
	for i, compileOpTrigger := range compileOpTriggers {
		compileOp := collectCompileOp(compileOpTrigger)
		if compileOp == nil {
			panic(fmt.Sprintf("the %d trigger did not call any gen.Compile internally", i))
		}
		compileOps = append(compileOps, compileOp)
	}
	generator := &generator{
		generatedTypes: map[reflect.Type]bool{},
		generatedFuncs: map[string]bool{},
	}
	source := ""
	for _, compileOp := range compileOps {
		_, oneSource := generator.gen(compileOp.template, compileOp.kv...)
		source += oneSource
	}
	logger.Debug("generated source", "source", source)
	source = prelog + source
	srcFileName := "/tmp/" + NewID().String() + ".go"
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
	err = os.Remove(srcFileName)
	if err != nil {
		logger.Error("failed to remove generated source", "srcFileName", srcFileName)
	}
}

func collectCompileOp(compileOpTrigger func()) (op *compileOp) {
	defer func() {
		recoved := recover()
		op, _ = recoved.(*compileOp)
	}()
	compileOpTrigger()
	return nil
}
