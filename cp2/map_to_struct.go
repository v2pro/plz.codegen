package cp2

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copyMapToStruct)
}

var copyMapToStruct = generic.DefineFunc("CopyMapToStruct(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"calcBindings", func(dstType, srcType reflect.Type) interface{} {
		bindings := []interface{}{}
		for i := 0; i < dstType.NumField(); i++ {
			dstField := dstType.Field(i)
			bindings = append(bindings, map[string]interface{}{
				"srcFieldName": dstField.Name,
				"srcFieldType": srcType.Elem(),
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
for key, elem := range src {
	switch key {
		{{ range $_, $binding := $bindings }}
			case "{{$binding.srcFieldName}}":
				{{$binding.cp}}(err, &dst.{{$binding.dstFieldName}}, elem)
		{{ end }}
	}
}`)
