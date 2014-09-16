package api

import (
	"github.com/mholt/binding"
	"net/http"
	"types"
)

type PaperForm struct {
	Title string
}

func (p *PaperForm) FieldMap() binding.FieldMap {
    return binding.FieldMap{
		&p.Title: binding.Field{
			Form:    "title",
			Required: true,
        },
    }
}

func PaperAdd(resp http.ResponseWriter, req *http.Request) {
	paperForm := &PaperForm{}
    if binding.Bind(req, paperForm).Handle(resp) {
        return
    }

	teacher := types.GetTeacherFromSession(req)
	p := types.Paper{}
	p.Teacher = teacher

	paperJson := types.PaperJson{}

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
		errorJson := types.AlreadyExists()
		paperJson = types.PaperJson{types.CreateStandardJsonErrorJson(req, errorJson), types.Paper{}}
		Render.JSON(resp, errorJson.Code, paperJson)
	} else {
		paperJson := types.PaperJson{types.CreateStandardJson(req), p}
		RenderJson(resp, paperJson)
	}
}

