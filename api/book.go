package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/mholt/binding"
	"net/http"
)

type BookForm struct {
	Title string
}

func (p *BookForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.Title: binding.Field{
			Form:     "title",
			Required: true,
		},
	}
}

func BookAdd(resp http.ResponseWriter, req *http.Request) {
	BookForm := &BookForm{}
	if binding.Bind(req, BookForm).Handle(resp) {
		return
	}

	teacher := models.GetTeacherFromSession(req)
	b := models.Book{}
	b.Teacher = teacher

	var alreadyExists bool

	query := `SELECT id, title, alreadyExists
			  FROM book_add($1, $2)`

	err := Db.QueryRow(query, BookForm.Title, teacher.Id).
		Scan(&b.Id, &b.Title, &alreadyExists)

	var dbError, noRows = checkQuery(err, resp, req, query)

	if dbError {
		return
	}

	if noRows || alreadyExists {
		errorJson := models.AlreadyExists()
		RenderErrorJson(resp, req, errorJson, models.Book{})
	} else {
		RenderJson(resp, req, b)
	}
}
