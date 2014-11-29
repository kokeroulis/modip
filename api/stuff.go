package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
)

func GetStuffList(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    d := models.ListAllStuff()
    RenderTemplate("stuff/list", helpers, resp, d)
}

func GetStuffCreate(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    var data interface{}
    RenderTemplate("stuff/create", helpers, resp, data)
}

func GetStuffEdit(resp http.ResponseWriter, req *http.Request) {
    var helpers []string

	id, _ := getId(req)

	t := models.Teacher{
		Id: id,
	}

	t.Load()
    RenderTemplate("stuff/edit", helpers, resp, t)
}

func StuffEdit(resp http.ResponseWriter, req *http.Request) {
	id, id2 := getId(req)

	t := models.Teacher{
		Id: id,
	}

	req.ParseForm()
	decoder := schema.NewDecoder()

	err := decoder.Decode(t, req.PostForm)

	if err != nil {
		panic(err)
	}

	t.Update()

	url := "/stuff/edit/" + id2
	http.Redirect(resp, req, url, http.StatusMovedPermanently)
}

