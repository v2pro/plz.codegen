package wombat

import (
	"strconv"
	"bytes"
	"net/http"
	"github.com/v2pro/plz"
	"fmt"
	_ "github.com/v2pro/wombat/cp_http"
)

func Example_bind_http_request() {
	body := "k1=v1&k2=v2"
	httpReq, _ := http.NewRequest("POST", "/url", bytes.NewBufferString(body))
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Content-Length", strconv.Itoa(len(body)))
	httpReq.Header.Add("Traceid", "1000")
	httpReq.ParseForm()

	type MyRequest struct {
		Url     string `http:"Url"`
		TraceId string `header:"Traceid"`
		K1      string `form:"k1"`
		K2      string `form:"k2"`
	}

	myReq := MyRequest{}
	plz.Copy(&myReq, httpReq)
	fmt.Println(myReq.Url)
	fmt.Println(myReq.TraceId)
	fmt.Println(myReq.K1)
	fmt.Println(myReq.K2)

	// Output:
	// /url
	// 1000
	// v1
	// v2
}
