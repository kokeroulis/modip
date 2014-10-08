package api

import (
	"io/ioutil"
	"net/http"
)

func Test(resp http.ResponseWriter, req *http.Request) {
	indexContents, err := ioutil.ReadFile("public/app/test.html")

	if err != nil {
		panic("Can't render test")
	}

	resp.Write(indexContents)
}
