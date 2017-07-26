package cpJson2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp2"
	"reflect"
	"github.com/v2pro/plz/lang/tagging"
)

func init() {
	cp2.Anything.ImportFunc(copyStructToJson)
	toJsonMap[reflect.Struct] = "CopyStructToJson"
}

var copyStructToJson = generic.DefineFunc(
	"CopyStructToJson(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp2.Anything).
	Generators(
	"calcBindings", func(dstType, srcType reflect.Type) interface{} {
		bindings := []interface{}{}
		tags := tagging.Get(srcType)
		for i := 0; i < srcType.NumField(); i++ {
			srcField := srcType.Field(i)
			jsonTag := tags.Fields[srcField.Name]["json"].Text()
			dstFieldName := srcField.Name
			if jsonTag != "" {
				dstFieldName = jsonTag
			}
			bindings = append(bindings, map[string]interface{}{
				"srcFieldName": srcField.Name,
				"srcFieldType": srcField.Type,
				"dstFieldName": dstFieldName,
				"dstFieldType": dstType,
			})
		}
		return bindings
	},
	"assignCp", func(binding map[string]interface{}, cpFuncName string) string {
		binding["cp"] = cpFuncName
		return ""
	}).
	Source(`
{{ $bindings := calcBindings .DT .ST }}
{{ range $_, $binding := $bindings}}
	{{ $cp := expand "CopyAnything" "DT" $binding.dstFieldType "ST" $binding.srcFieldType }}
	{{ assignCp $binding $cp }}
{{ end }}
dst.WriteObjectStart()
{{ range $i, $binding := $bindings }}
	{{ if ne $i 0 }}
		dst.WriteMore()
	{{ end }}
	dst.WriteObjectField("{{$binding.dstFieldName}}")
	{{$binding.cp}}(err, dst, src.{{$binding.srcFieldName}})
{{ end }}
dst.WriteObjectEnd()`)
