package api

import (
	"encoding/gob"
	"github.com/kokeroulis/modip/models"
	"github.com/kokeroulis/modip/models/forms"
	"github.com/mholt/binding"
	"github.com/gorilla/schema"
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

func GetTeacherCreateReport(resp http.ResponseWriter, req *http.Request) {
    helpers := []string{
        "templates/teacher/anagnorish_epistomonikou_ergou.tmpl",
        "templates/teacher/eureunitika_programmata_kai_erga.tmpl",
        "templates/teacher/sundesi_me_tin_koinonia.tmpl",
        "templates/teacher/arithmos_dimosieuseon.tmpl",
    }

	t := forms.TeacherCreateReportForm{
		TeacherId: models.GetTeacherFromSession(req).Id,
	}

	t.Load()

	RenderTemplate("teacher/report", helpers, resp, t)
}

func TeacherCreateReport1(resp http.ResponseWriter, req *http.Request) {
	teacherId := models.GetTeacherFromSession(req).Id

	req.ParseForm()

	form := &forms.TeacherCreateReportFormEntry1{}
	parseTeacherCreateReportFormEntry1(form, req)

	form.Create(teacherId)

    http.Redirect(resp, req, "/teacher/report", http.StatusMovedPermanently)
}

func TeacherCreateReport(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	formNumber, id := getId(req)
	teacherId := models.GetTeacherFromSession(req).Id

	decoder := schema.NewDecoder()

	var err error
	switch formNumber {
	case 2:
		form := &forms.TeacherCreateReportFormEntry2{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId)
	case 3:
		form := &forms.TeacherCreateReportFormEntry3{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId)
	case 4:
		form := &forms.TeacherCreateReportFormEntry4{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId)
	case 5:
		form := &forms.TeacherCreateReportFormEntry5{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId)
	default:
		unknownErr := "Unknown form: " + id
		panic(unknownErr)
	}

	if err != nil {
		panic(err)
	}
    http.Redirect(resp, req, "/teacher/report", http.StatusMovedPermanently)
}

func GetTeacherCreateReport1Edit(resp http.ResponseWriter, req *http.Request) {
    entryId, _ := getId(req)
	var helpers []string
    formsList := forms.ListAllTeacherCreateReportForm1(models.GetTeacherFromSession(req).Id)

    var data interface{}

    for _, it := range formsList {
        if it.Id == entryId {
            data = it
        }
    }

    RenderTemplate("teacher/report_edit", helpers, resp, data);
}

func TeacherCreateReport1Edit(resp http.ResponseWriter, req *http.Request) {
	form := &forms.TeacherCreateReportFormEntry1{}
	parseTeacherCreateReportFormEntry1(form, req)

	id, _ := getId(req)
	form.Id = id
	form.Update()
    http.Redirect(resp, req, "/teacher/report", http.StatusMovedPermanently)
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

