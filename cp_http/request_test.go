package cp_http

import (
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/require"
	"net/url"
	"strconv"
	"github.com/v2pro/plz"
)

func Test_req_method(t *testing.T) {
	should := require.New(t)

	type TestObject struct {
		Field string `http:"Method"`
	}

	req, _ := http.NewRequest("GET", "/", nil)
	obj := TestObject{}
	should.Nil(plz.Copy(&obj, req))
	should.Equal("GET", obj.Field)
}

func createFormRequest(kv ...string) *http.Request {
	data := url.Values{}
	for i := 0; i < len(kv); i += 2 {
		data.Add(kv[i], kv[i+1])
	}
	body := data.Encode()
	req, err := http.NewRequest("POST", "/", bytes.NewBufferString(body))
	if err != nil {
		panic(err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))
	return req
}
