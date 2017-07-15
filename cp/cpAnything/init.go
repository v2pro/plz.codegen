package cpAnything

import (
	"fmt"
	"github.com/v2pro/plz"
	"github.com/v2pro/plz/logging"
	"github.com/v2pro/plz/util"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

var Dispatchers = []func(dstType, srcType reflect.Type) string{}
var logger = plz.LoggerOf("package", "cpAnything")

func init() {
	util.GenCopy = Gen
	logging.Providers = append(logging.Providers, func(loggerKv []interface{}) logging.Logger {
		for i := 0; i < len(loggerKv); i += 2 {
			key := loggerKv[i].(string)
			if key == "package" && "cpAnything" == loggerKv[i+1] {
				return logging.NewStderrLogger(loggerKv, logging.LEVEL_DEBUG)
			}
		}
		return nil
	})
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName: "cpAnything",
	GenMap: map[string]interface{}{
		"dispatch": genDispatch,
	},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName:     `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Declarations: "var cpDynamically func(interface{}, interface{}) error",
	Source: `
{{ $tmpl := dispatch .DT .ST }}
{{ $cp := gen $tmpl "DT" .DT "ST" .ST }}
{{ $cp.Source }}

func Exported_{{ .funcName }}(
	cp func(interface{}, interface{}) error,
	dst interface{},
	src interface{}) (err error) {
	// end of signature
	cpDynamically = cp
	{{ .funcName }}(
		&err,
		{{ cast "dst" .DT }},
		{{ cast "src" .ST }})
	return
}
`,
}

func genDispatch(dstType, srcType reflect.Type) string {
	template := dispatch(dstType, srcType)
	logger.Info("dispatch result", "dstType", dstType, "srcType", srcType, "template", template)
	return template
}

func dispatch(dstType, srcType reflect.Type) string {
	logger.Debug("dispatch", "dstType", dstType, "srcType", srcType)
	for _, dispatcher := range Dispatchers {
		tmpl := dispatcher(dstType, srcType)
		if tmpl != "" {
			return tmpl
		}
	}
	if dstType.Kind() != reflect.Ptr && dstType.Kind() != reflect.Map && dstType.Kind() != reflect.Interface {
		panic("destination type is not writable: " + dstType.String())
	}
	if srcType.Kind() == reflect.Ptr {
		return "cpFromPtr"
	}
	if srcType.Kind() == reflect.Interface && srcType.NumMethod() == 0 {
		return "cpFromInterface"
	}
	if dstType.Kind() == reflect.Map &&
		srcType.Kind() == reflect.Map {
		return "cpMapToMap"
	}
	if dstType.Kind() == reflect.Map &&
		srcType.Kind() == reflect.Struct {
		return "cpStructToMap"
	}
	if isPtrPtr(dstType) {
		return "cpIntoPtr"
	}
	if dstType.Kind() == reflect.Ptr {
		if isSimpleValue(dstType.Elem()) && dstType.Elem().Kind() == srcType.Kind() {
			return "cpSimpleValue"
		}
		if dstType.Elem().Kind() == reflect.Interface {
			return "cpIntoInterface"
		}
		if dstType.Elem().Kind() == reflect.Struct &&
			srcType.Kind() == reflect.Struct {
			return "cpStructToStruct"
		}
		if dstType.Elem().Kind() == reflect.Struct &&
			srcType.Kind() == reflect.Map {
			return "cpMapToStruct"
		}
		if dstType.Elem().Kind() == reflect.Slice &&
			srcType.Kind() == reflect.Slice {
			return "cpSliceToSlice"
		}
		if dstType.Elem().Kind() == reflect.Array &&
			srcType.Kind() == reflect.Array {
			return "cpArrayToArray"
		}
		if dstType.Elem().Kind() == reflect.Slice &&
			srcType.Kind() == reflect.Array {
			return "cpSliceToSlice"
		}
		if dstType.Elem().Kind() == reflect.Array &&
			srcType.Kind() == reflect.Slice {
			return "cpSliceToSlice"
		}
	}
	panic(fmt.Sprintf("not implemented copy: from %v to %v", srcType, dstType))
}

func isPtrPtr(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr && (
		typ.Elem().Kind() == reflect.Ptr ||
			typ.Elem().Kind() == reflect.Map)
}

func isSimpleValue(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:
		return true
	}
	return false
}

// Gen generates a instance of F
func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	f := funcObj.(func(func(interface{}, interface{}) error, interface{}, interface{}) error)
	return func(dst, src interface{}) error {
		return f(util.Copy, dst, src)
	}
}
