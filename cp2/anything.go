package cp2

import (
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/plz/util"
	"reflect"
)

func init() {
	util.GenCopy = func(dstType reflect.Type, srcType reflect.Type) func(interface{}, interface{}) error {
		var ct func(interface{}, interface{}) error
		funcObj := generic.Expand(AnythingForPlz,
			"DT", dstType,
			"ST", srcType,
			"CT", reflect.TypeOf(ct))
		f := funcObj.(func(func(interface{}, interface{}) error, interface{}, interface{}) error)
		return func(dst, src interface{}) error {
			return f(util.Copy, dst, src)
		}
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
		switch srcType.Elem().Kind() {
		case reflect.Ptr, reflect.Map, reflect.Slice:
			return "CopyFromPtrPtr"
		}
	}
	if srcType.Kind() == reflect.Interface && srcType.NumMethod() == 0 {
		return "CopyFromInterface"
	}
	if dstType.Kind() == reflect.Map &&
		srcType.Kind() == reflect.Map {
		return "CopyMapToMap"
	}
	if dstType.Kind() == reflect.Map &&
		srcType.Kind() == reflect.Struct {
		return "CopyStructToMap"
	}
	if dstType.Kind() == reflect.Ptr {
		if dstType.Elem().Kind() == reflect.Ptr || dstType.Elem().Kind() == reflect.Map {
			return "CopyIntoPtr"
		}
		if srcType.Kind() == reflect.Slice && dstType.Elem().Kind() == reflect.Array {
			return "CopySliceToArray"
		}
		if srcType.Kind() == reflect.Array && dstType.Elem().Kind() == reflect.Slice {
			return "CopySliceToSlice"
		}
		if srcType.Kind() == reflect.Map && dstType.Elem().Kind() == reflect.Struct {
			return "CopyMapToStruct"
		}
		if dstType.Elem().Kind() == srcType.Kind() {
			switch dstType.Elem().Kind() {
			case reflect.Array:
				return "CopyArrayToArray"
			case reflect.Slice:
				return "CopySliceToSlice"
			case reflect.Struct:
				return "CopyStructToStruct"
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Bool, reflect.String, reflect.Float32, reflect.Float64:
				return "CopySimpleValue"
			}
		}
	}
	if srcType.Kind() == reflect.Ptr {
		return "CopyFromPtr"
	}
	panic("do not know how to copy " + srcType.String() + " to " + dstType.String())
}

var AnythingForPlz = generic.DefineFunc("CopyAnythingForPlz(theCopyDynamically CT, dst interface{}, src interface{}) error").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	Param("CT", "type of copy dynamically function").
	ImportFunc(Anything).
	Declare("var copyDynamically func(interface{}, interface{}) error").
	Source(`
{{ $cp := expand "CopyAnything" "DT" .DT "ST" .ST }}
copyDynamically = theCopyDynamically
var err error
{{$cp}}(&err, dst.({{.DT|name}}), src.({{.ST|name}}))
return err`)
