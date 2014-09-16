package api

import (
	"github.com/mholt/binding"
	"net/http"
	"github.com/kokeroulis/modip/types"
)

type BookForm struct {
	Title string
}

func (p *BookForm) FieldMap() binding.FieldMap {
    return binding.FieldMap{
		&p.Title: binding.Field{
			Form:    "title",
			Required: true,
        },
    }
}

func BookAdd(resp http.ResponseWriter, req *http.Request) {
	BookForm := &BookForm{}
    if binding.Bind(req, BookForm).Handle(resp) {
        return
    }

	teacher := types.GetTeacherFromSession(req)
	b := types.Book{}
	b.Teacher = teacher

	bookJson := types.BookJson{}

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
		errorJson := types.AlreadyExists()
		bookJson = types.BookJson{types.CreateStandardJsonErrorJson(req, errorJson), types.Book{}}
		Render.JSON(resp, errorJson.Code, bookJson)
	} else {
		bookJson := types.BookJson{types.CreateStandardJson(req), b}
		RenderJson(resp, bookJson)
	}
}

