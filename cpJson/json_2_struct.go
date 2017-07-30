package cpJson

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/cp"
	"reflect"
	"github.com/v2pro/plz/lang/tagging"
)

func init() {
	cp.Anything.ImportFunc(copyJsonToStruct)
	fromJsonMap[reflect.Struct] = "CopyJsonToStruct"
}

var copyJsonToStruct = generic.DefineFunc(
	"CopyJsonToStruct(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(cp.Anything).
	Generators(
	"calcBindings", func(dstType, srcType reflect.Type) interface{} {
		bindings := []interface{}{}
		tags := tagging.Get(dstType)
		for i := 0; i < dstType.NumField(); i++ {
			dstField := dstType.Field(i)
			srcFieldName := dstField.Name
			jsonTag := tags.Fields[dstField.Name]["json"].Text()
			if jsonTag != "" {
				srcFieldName = jsonTag
			}
			bindings = append(bindings, map[string]interface{}{
				"srcFieldName": srcFieldName,
				"srcFieldType": srcType,
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
src.ReadObjectCB(func(iter *jsoniter.Iterator, field string) bool {
	switch field {
		{{ range $_, $binding := $bindings }}
			case "{{ $binding.srcFieldName }}":
				{{$binding.cp}}(err, &dst.{{$binding.dstFieldName}}, iter)
		{{ end }}
		default:
			iter.Skip()
	}
	return true
})`)
