package cp_http

import (
	"github.com/v2pro/plz/lang"
	"unsafe"
	"net/http"
	"reflect"
)

func newRequestAccessor(tagName string) lang.Accessor {
	index := 0
	fields := []*requestField{
		newStringRequestField(&index, "Method", func(req *http.Request) string {
			return req.Method
		}),
	}
	return &requestAccessor{
		lang.NoopAccessor{tagName, "requestAccessor"},
		fields,
	}
}

type requestAccessor struct {
	lang.NoopAccessor
	fields []*requestField
}

func (accessor *requestAccessor) ReadOnly() bool {
	return false
}

func (accessor *requestAccessor) Kind() lang.Kind {
	return lang.Struct
}

func (accessor *requestAccessor) GoString() string {
	return "jsoncp.streamAccessor"
}

func (accessor *requestAccessor) NumField() int {
	return len(accessor.fields)
}

func (accessor *requestAccessor) Field(index int) lang.StructField {
	return accessor.fields[index]
}

func (accessor *requestAccessor) IterateArray(ptr unsafe.Pointer, cb func(index int, elem unsafe.Pointer) bool) {
	req := (*http.Request)(ptr)
	for i := 0; i < len(accessor.fields); i++ {
		if !accessor.fields[i].callCb(req, cb) {
			return
		}
	}
}

type requestField struct {
	index    int
	name     string
	accessor lang.Accessor
	callCb   func(req *http.Request, cb func(index int, elem unsafe.Pointer) bool) bool
}

func newStringRequestField(index *int, name string, getValue func(req *http.Request) string) *requestField {
	currentIndex := *index
	*index = currentIndex + 1
	return &requestField{
		index:    currentIndex,
		name:     name,
		accessor: lang.AccessorOf(reflect.TypeOf(""), ""),
		callCb: func(req *http.Request, cb func(index int, elem unsafe.Pointer) bool) bool {
			return cb(currentIndex, lang.AddressOf(getValue(req)))
		},
	}
}

func (field *requestField) Index() int {
	return field.index
}
func (field *requestField) Name() string {
	return field.name
}
func (field *requestField) Accessor() lang.Accessor {
	return field.accessor
}
func (field *requestField) Tags() map[string]interface{} {
	return map[string]interface{}{}
}
