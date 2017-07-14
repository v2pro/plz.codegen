package wombat

import (
	// register fp functions to plz
	_ "github.com/v2pro/wombat/fp"
	"github.com/v2pro/wombat/gen"
)

// CompilePlugin export from gen
func CompilePlugin(soFileName string, compileOpTriggers ...func()) {
	gen.CompilePlugin(soFileName, compileOpTriggers...)
}

// LoadPlugin export from gen
func LoadPlugin(soFileName string) {
	gen.LoadPlugin(soFileName)
}

// DisableDynamicCompilation export from gen
func DisableDynamicCompilation() {
	gen.DisableDynamicCompilation()
}
