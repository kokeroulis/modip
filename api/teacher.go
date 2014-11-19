package api

import (
	"encoding/gob"
	"github.com/kokeroulis/modip/models"
	"github.com/kokeroulis/modip/models/forms"
	"github.com/mholt/binding"
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func GetTeacherCreateReport(resp http.ResponseWriter, req *http.Request) {
    helpers := []string{
        "templates/teacher/anagnorish_epistomonikou_ergou.tmpl",
        "templates/teacher/eureunitika_programmata_kai_erga.tmpl",
        "templates/teacher/sundesi_me_tin_koinonia.tmpl",
    }

	RenderTemplate("teacher/report", helpers, resp,
					forms.ListAllTeacherCreateReportForm1(models.GetTeacherFromSession(req).Id))
}

func TeacherCreateReport1(resp http.ResponseWriter, req *http.Request) {
	teacherId := models.GetTeacherFromSession(req).Id

	req.ParseForm()

	form := forms.TeacherCreateReportFormEntry1{}

	decoder := schema.NewDecoder()
	err := decoder.Decode(form, req.PostForm)
	form.Create(teacherId)

	if err != nil {
		panic(err)
	}
}

func TeacherCreateReport(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	vars := mux.Vars(req)
	id, ok := vars["id"]

	formNumber, paramErr := strconv.Atoi(id)

	if paramErr != nil || !ok {
		panic(paramErr)
	}

//	teacherId := models.GetTeacherFromSession(req).Id

	switch formNumber {
	case 2:
		//form := &forms.TeacherCreateReportFormEntry2}
		//err = decoder.Decode(form, req.PostForm)
		//form.Create(teacherId)
	default:
		unknownErr := "Unknown form: " + id
		panic(unknownErr)
	}


}

func GetTeacherCreateReport1Edit(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}

	RenderTemplate("teacher/report_edit", helpers, resp, data)
}

func TeacherCreateReport1Edit(resp http.ResponseWriter, req *http.Request) {
	form := &forms.TeacherCreateReportFormEntry1{}
	parseTeacherCreateReportFormEntry1(form, req)

	form.Id = getId(req)
	form.Update()
}

func parseTeacherCreateReportFormEntry1(form *forms.TeacherCreateReportFormEntry1,
										req *http.Request) {
	req.ParseForm()
	decoder := schema.NewDecoder()
	err := decoder.Decode(form, req.PostForm)

	if err != nil {
		panic(err)
	}
}

func getId(req *http.Request) int {
	vars := mux.Vars(req)
	id, ok := vars["id"]

	res, err := strconv.Atoi(id)

	if err != nil || !ok {
		panic(err)
	}

	return res
}
