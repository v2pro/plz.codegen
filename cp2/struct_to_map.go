package cp2

import (
	"github.com/v2pro/wombat/generic"
	"reflect"
)

func init() {
	Anything.ImportFunc(copyStructToMap)
}

var copyStructToMap = generic.DefineFunc("CopyStructToMap(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Generators(
	"calcBindings", func(dstType, srcType reflect.Type) interface{} {
		bindings := []interface{}{}
		for i := 0; i < srcType.NumField(); i++ {
			srcField := srcType.Field(i)
			bindings = append(bindings, map[string]interface{}{
				"srcFieldName": srcField.Name,
				"srcFieldType": srcField.Type,
				"dstFieldName": srcField.Name,
				"dstFieldType": reflect.PtrTo(dstType.Elem()),
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
var existingElem {{.DT|elem|name}}
var found bool
{{ range $_, $binding := $bindings }}
	existingElem, found = dst["{{$binding.dstFieldName}}"]
	if found {
		{{$binding.cp}}(err, &existingElem, src.{{$binding.srcFieldName}})
		dst["{{$binding.dstFieldName}}"] = existingElem
	} else {
		newElem := new({{$binding.dstFieldType|elem|name}})
		{{$binding.cp}}(err, newElem, src.{{$binding.srcFieldName}})
		dst["{{ $binding.dstFieldName }}"] = *newElem
	}
{{ end }}`)
