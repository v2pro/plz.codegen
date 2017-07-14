package cpStatically

import (
	"github.com/v2pro/plz"
	"github.com/v2pro/plz/logging"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

var logger = plz.LoggerOf("package", "cpStatically")

func init() {
	logging.Providers = append(logging.Providers, func(loggerKv []interface{}) logging.Logger {
		for i := 0; i < len(loggerKv); i += 2 {
			key := loggerKv[i].(string)
			if key == "package" && "cpStatically" == loggerKv[i+1] {
				return logging.NewStderrLogger(loggerKv, logging.LEVEL_DEBUG)
			}
		}
		return nil
	})
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{},
	FuncMap: map[string]interface{}{
		"dispatch": dispatch,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $tmpl := dispatch .DT .ST }}
{{ $cp := gen $tmpl "DT" .DT "ST" .ST }}
{{ $cp.Source }}

func Exported_{{ .funcName }}(
	dst interface{},
	src interface{}) error {
	// end of signature
	return {{ .funcName }}(
		{{ cast "dst" .DT }},
		{{ cast "src" .ST }})
}
`,
}

func dispatch(dstType, srcType reflect.Type) string {
	template := doDispatch(dstType, srcType)
	logger.Info("dispatch result", "dstType", dstType, "srcType", srcType, "template", template)
	return template
}

func doDispatch(dstType, srcType reflect.Type) string {
	if dstType.Kind() != reflect.Ptr && dstType.Kind() != reflect.Map {
		panic("destination type is not writable: " + dstType.String())
	}
	if srcType.Kind() == reflect.Ptr {
		return "cpFromPtr"
	} else {
		if dstType.Kind() == reflect.Ptr {
			if isDirectPtr(dstType) {
				if isSimpleValue(dstType.Elem()) {
					return "cpSimpleValue"
				} else if dstType.Elem().Kind() == reflect.Struct && srcType.Kind() == reflect.Struct {
					return "cpStructToStruct"
				} else if dstType.Elem().Kind() == reflect.Slice && srcType.Kind() == reflect.Slice {
					return "cpSliceToSlice"
				} else if dstType.Elem().Kind() == reflect.Array && srcType.Kind() == reflect.Array {
					return "cpArrayToArray"
				} else {
					panic("not implemented")
				}
			} else {
				return "cpIntoPtr"
			}
		} else if dstType.Kind() == reflect.Map && srcType.Kind() == reflect.Map {
			return "cpMapToMap"
		} else {
			panic("not implemented")
		}
	}
}

func isDirectPtr(typ reflect.Type) bool {
	if typ.Kind() != reflect.Ptr {
		return false
	}
	return typ.Elem().Kind() != reflect.Ptr
}

func isSimpleValue(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func Gen(dstType, srcType reflect.Type) func(interface{}, interface{}) error {
	funcObj := gen.Compile(F, "DT", dstType, "ST", srcType)
	return funcObj.(func(interface{}, interface{}) error)
}
