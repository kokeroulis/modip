package api

import (
	"io/ioutil"
	"net/http"
)

func Index(resp http.ResponseWriter, req *http.Request) {
	indexContents, err := ioutil.ReadFile("public/index.html")

	if err != nil {
		panic("Can't render index")
	}

	resp.Write(indexContents)
}
