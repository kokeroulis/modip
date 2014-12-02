package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
	"strconv"
)


func GetAkademicYearList(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("akademic_year/list", helpers, resp,models.ListAkademicYears())
}

func GetAkademicYearCreate(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}

	RenderTemplate("akademic_year/create", helpers, resp, data)
}

func AkademicYearCreate(resp http.ResponseWriter, req *http.Request) {
	a := &models.AkademicYear{}

	req.ParseForm()
	decoder := schema.NewDecoder()

	err := decoder.Decode(a, req.PostForm)

	if err != nil {
		panic(err)
	}

	a.Create()

	url := "/akademic/year/list"
	http.Redirect(resp, req, url, http.StatusMovedPermanently)
}

func GetAkademicYearEdit(resp http.ResponseWriter, req *http.Request) {
    var helpers []string

	id, _ := getId(req)

	a := models.AkademicYear{
		Id: id,
	}

	a.Load()
    RenderTemplate("akademic_year/edit", helpers, resp, a)
}

func AkademicYearEdit(resp http.ResponseWriter, req *http.Request) {
	a := &models.AkademicYear{}

	req.ParseForm()
	decoder := schema.NewDecoder()

	err := decoder.Decode(a, req.PostForm)

	if err != nil {
		panic(err)
	}

	a.Update()

	url := "/akademic/year/edit/" + strconv.Itoa(a.Id)
	http.Redirect(resp, req, url, http.StatusMovedPermanently)
}

