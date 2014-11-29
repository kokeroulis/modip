package api

import (
	"github.com/kokeroulis/modip/models"
	_"github.com/gorilla/schema"
	"net/http"
)

func GetStuffList(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    RenderTemplate("stuff/list", helpers, resp, models.ListAllStuff())
}

func GetStuffCreate(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    var data interface{}
    RenderTemplate("stuff/create", helpers, resp, data)
}

func GetStuffEdit(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    var data interface{}
    RenderTemplate("stuff/edit", helpers, resp, data)
}
