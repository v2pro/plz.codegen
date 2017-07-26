package cp2

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copyStructToStruct)
}

var copyStructToStruct = generic.DefineFunc("CopyStructToStruct(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"calcBindings", func(dstType, srcType reflect.Type) interface{} {
		bindings := []interface{}{}
		for i := 0; i < dstType.NumField(); i++ {
			dstField := dstType.Field(i)
			srcField, srcFieldFound := srcType.FieldByName(dstField.Name)
			if !srcFieldFound {
				continue
			}
			bindings = append(bindings, map[string]interface{}{
				"srcFieldName": srcField.Name,
				"srcFieldType": srcField.Type,
				"dstFieldName": dstField.Name,
				"dstFieldType": reflect.PtrTo(dstField.Type),
			})
		}
		return bindings
	},
	"assignCp", func(binding map[string]interface{}, cpFuncName string) string {
		binding["cp"] = cpFuncName
		return ""
	}).
	Source(`
{{ $bindings := calcBindings (.DT|elem) .ST }}
{{ range $_, $binding := $bindings}}
	{{ $cp := expand "CopyAnything" "DT" $binding.dstFieldType "ST" $binding.srcFieldType }}
	{{ assignCp $binding $cp }}
{{ end }}
{{ range $_, $binding := $bindings }}
	{{$binding.cp}}(err, &dst.{{$binding.dstFieldName}}, src.{{$binding.srcFieldName}})
{{ end }}`)
