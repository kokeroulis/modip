package api

import (
	"encoding/gob"
	"github.com/kokeroulis/modip/models"
	"github.com/mholt/binding"
	"net/http"
)

type TeacherForm struct {
	Username string
	Password string
}

func (t *TeacherForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.Username: binding.Field{
			Form:     "username",
			Required: true,
		},
		&t.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

func init() {
	// later we will try to save to our
	// session a struct. Structs are complex
	// models which aren't supported by default
	// from the god package so we have to
	// register it manually
	gob.Register(&models.Teacher{})
}

func TeacherLogin(resp http.ResponseWriter, req *http.Request) {
	teacherForm := &TeacherForm{}
	if binding.Bind(req, teacherForm).Handle(resp) {
		return
	}

	t := models.Teacher{}
	auth := t.Login(teacherForm.Username, teacherForm.Password)

	if auth {
		session := models.GetSession(req)
		session.Values["teacher"] = &t
		session.Save(req, resp)
		RenderJson(resp, req, t)
	} else {
		errorJson := models.AuthFailed()
		RenderErrorJson(resp, req, errorJson, t)
	}
}

func GetTeacherReport(resp http.ResponseWriter, req *http.Request) {
    helpers := []string{
        "templates/teacher/anagnorish_epistomonikou_ergou.tmpl",
        "templates/teacher/eureunitika_programmata_kai_erga.tmpl",
        "templates/teacher/sundesi_me_tin_koinonia.tmpl",
    }
	var data interface{}
	RenderTemplate("teacher/report", helpers, resp, data)
}
