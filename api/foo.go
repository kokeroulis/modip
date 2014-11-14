package api

import "net/http"

func Foo(resp http.ResponseWriter, req *http.Request) {
	var data interface{}

	RenderTemplate("foo", resp, data)
}
