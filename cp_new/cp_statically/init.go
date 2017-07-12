package cp_statically

import (
	"reflect"
	"github.com/v2pro/wombat/gen"
	"github.com/v2pro/plz"
	"github.com/v2pro/plz/logging"
)

var logger = plz.LoggerOf("package", "cp_statically")

func init() {
	logging.Providers = append(logging.Providers, func(loggerKv []interface{}) logging.Logger {
		for i := 0; i < len(loggerKv); i += 2 {
			key := loggerKv[i].(string)
			if key == "package" && "cp_statically" == loggerKv[i+1] {
				return logging.NewStderrLogger(loggerKv, logging.LEVEL_DEBUG)
			}
		}
		return nil
	})
}

var F = &gen.FuncTemplate{
	Dependencies: map[string]*gen.FuncTemplate{
	},
	FuncMap: map[string]interface{}{
		"dispatch": func_dispatch,
	},
	Variables: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `Copy_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $tmpl := dispatch .DT .ST }}
{{ $cp := gen $tmpl "DT" .DT "ST" .ST }}
{{ $cp.Source }}
`,
}

func func_dispatch(dstType, srcType reflect.Type) string {
	template := _func_dispatch(dstType, srcType)
	logger.Info("dispatch result", "dstType", dstType, "srcType", srcType, "template", template)
	return template
}

func _func_dispatch(dstType, srcType reflect.Type) string {
	if dstType.Kind() != reflect.Ptr && dstType.Kind() != reflect.Map {
		panic("destination type is not writable: " + dstType.String())
	}
	if srcType.Kind() == reflect.Ptr && srcType.Elem().Kind() != reflect.Struct {
		return "cp_from_ptr"
	} else {
		if dstType.Kind() == reflect.Ptr {
			if isDirectPtr(dstType) {
				if isSimpleValue(dstType.Elem()) {
					return "cp_simple_value"
				} else if dstType.Elem().Kind() == reflect.Struct {
					if srcType.Kind() == reflect.Struct {
						return "cp_struct_to_struct"
					} else if srcType.Kind() == reflect.Ptr && srcType.Elem().Kind() == reflect.Struct {
						return "cp_ptr_struct_to_struct"
					} else {
						panic("not implemented")
					}
				} else {
					panic("not implemented")
				}
			} else {
				return "cp_into_ptr"
			}
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
