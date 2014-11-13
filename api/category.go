package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
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

type CategorySaveForm struct {
	CategoryId int
	Data       string
}

func (c *CategorySaveForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&c.CategoryId: binding.Field{
			Form:     "categoryid",
			Required: true,
		},
		&c.Data: binding.Field{
			Form:     "data",
			Required: true,
		},
	}
}

func CategorySave(resp http.ResponseWriter, req *http.Request) {
	categorySaveForm := &CategorySaveForm{}
	if binding.Bind(req, categorySaveForm).Handle(resp) {
		return
	}

	c := models.Category{
		Id: categorySaveForm.CategoryId,
		Data: categorySaveForm.Data,
	}

	c.Save(models.GetTeacherFromSession(req).Id)

	RenderJson(resp, req, c)
}
