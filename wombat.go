package wombat

import (
	// register fp functions to plz
	_ "github.com/v2pro/wombat/fp"
	"github.com/v2pro/wombat/gen"
	"reflect"
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

// Expand export from gen
func Expand(template *gen.FuncTemplate, templateArgs ...interface{}) {
	gen.Expand(template, templateArgs...)
}

// CompilePlugin export from gen
func CompilePlugin(soFileName string) {
	gen.CompilePlugin(soFileName)
}

// LoadPlugin export from gen
func LoadPlugin(soFileName string) {
	gen.LoadPlugin(soFileName)
}

// DisableDynamicCompilation export from gen
func DisableDynamicCompilation() {
	gen.DisableDynamicCompilation()
}
