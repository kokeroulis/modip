package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CategoryList(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]

	categoryId, err := strconv.Atoi(id)

	if categoryId <= 0 || err != nil || !ok {
		RenderJson(resp, req, models.ListAllCategories())
		return
	}

	c := models.Category{
		Id: categoryId,
	}

	c.Load()
	RenderJson(resp, req, c)
}
