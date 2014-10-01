package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/mholt/binding"
	"net/http"
)

type PaperForm struct {
	Title string
}

func (p *PaperForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.Title: binding.Field{
			Form:     "title",
			Required: true,
		},
	}
}

func PaperAdd(resp http.ResponseWriter, req *http.Request) {
	paperForm := &PaperForm{}
	if binding.Bind(req, paperForm).Handle(resp) {
		return
	}

	teacher := models.GetTeacherFromSession(req)
	p := models.Paper{}
	p.Teacher = teacher


	var alreadyExists bool

	query := `SELECT id, title, alreadyExists
			  FROM paper_add($1, $2)`

	err := Db.QueryRow(query, paperForm.Title, teacher.Id).
		Scan(&p.Id, &p.Title, &alreadyExists)

	var dbError, noRows = checkQuery(err, resp, req, query)

	if dbError {
		return
	}

	if noRows || alreadyExists {
		errorJson := models.AlreadyExists()
		RenderErrorJson(resp, req, errorJson, models.Paper{})
	} else {
		RenderJson2(resp, req, p)
	}
}
