package cp2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/plz/util"
	"reflect"
)

func init() {
	util.GenCopy = func(dstType reflect.Type, srcType reflect.Type) func(interface{}, interface{}) error {
		funcObj := generic.Expand(AnythingForPlz, "DT", dstType, "ST", srcType)
		f := funcObj.(func(interface{}, interface{}) error)
		return f
	}
}

var Anything = generic.DefineFunc("CopyAnything(err *error, dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	Generators("dispatch", dispatch).
	Source(`
{{ $tmpl := dispatch .DT .ST }}
{{ $cp := expand $tmpl "DT" .DT "ST" .ST }}
{{$cp}}(err, dst, src)`)

func dispatch(dstType reflect.Type, srcType reflect.Type) string {
	if srcType.Kind() == reflect.Ptr {
		return "CopyFromPtr"
	}
	if dstType.Kind() == reflect.Ptr {
		switch dstType.Elem().Kind() {
		case reflect.Ptr:
			return "CopyIntoPtr"
		case reflect.Int, reflect.Int8:
			if srcType.Kind() == dstType.Elem().Kind() {
				return "CopySimpleValue"
			}
		}
	}
	panic("do not know how to copy " + srcType.String() + " to " + dstType.String())
}

var AnythingForPlz = generic.DefineFunc("CopyAnythingForPlz(dst interface{}, src interface{}) error").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	ImportFunc(Anything).
	Source(`
{{ $cp := expand "CopyAnything" "DT" .DT "ST" .ST }}
var err error
{{$cp}}(&err, dst.({{.DT|name}}), src.({{.ST|name}}))
return err`)