package cp_http

import (
	"unsafe"
	"net/http"
	"github.com/v2pro/plz/lang/tagging"
)

type requestWrapper struct {
	Req *http.Request `http:"-"`
}

func (wrapper *requestWrapper) DefineTags() tagging.Tags {
	return tagging.D(
		tagging.S(),
		tagging.F(tagging.VirtualField("Method"), "http", tagging.V(
			"mapValue", func(ptr unsafe.Pointer) interface{} {
				obj := (*requestWrapper)(ptr)
				return &obj.Req.Method
			},
		)),
		tagging.F(tagging.VirtualField("Url"), "http", tagging.V(
			"mapValue", func(ptr unsafe.Pointer) interface{} {
				obj := (*requestWrapper)(ptr)
				return obj.Req.URL.String()
			},
		)),
		tagging.F(tagging.VirtualField("Header"), "http", tagging.V(
			"mapValue", func(ptr unsafe.Pointer) interface{} {
				obj := (*requestWrapper)(ptr)
				return &obj.Req.Header
			},
		)),
		tagging.F(tagging.VirtualField("Form"), "http", tagging.V(
			"mapValue", func(ptr unsafe.Pointer) interface{} {
				obj := (*requestWrapper)(ptr)
				return &obj.Req.Form
			},
		)),
	)
}
