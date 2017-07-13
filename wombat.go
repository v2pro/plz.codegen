package wombat

import (
	_ "github.com/v2pro/wombat/fp"
	"github.com/v2pro/wombat/gen"
)

func CompilePlugin(soFileName string, compileOpTriggers ...func()) {
	gen.CompilePlugin(soFileName, compileOpTriggers...)
}

func LoadPlugin(soFileName string) {
	gen.LoadPlugin(soFileName)
}

func DisableDynamicCompilation() {
	gen.DisableDynamicCompilation()
}
