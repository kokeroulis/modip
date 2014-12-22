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

func GetTeacherCreateReport(resp http.ResponseWriter, req *http.Request) {
    helpers := []string{
        "templates/teacher/arithmos_dimosieuseon.tmpl",
        "templates/teacher/ereunitikes_ypodomes.tmpl",
        "templates/teacher/anagnorish_epistomonikou_ergou.tmpl",
        "templates/teacher/eureunitika_programmata_kai_erga.tmpl",
        "templates/teacher/sundesi_me_tin_koinonia.tmpl",
    }
    akademicYearId, _ := getAkademicYearId(req)

	t := forms.TeacherCreateReportForm{
		TeacherId: models.GetTeacherFromSession(req).Id,
        AkademicYearId: akademicYearId,
	}

	t.Load()

    //TODO FIXME
    //Why are we doing this?
    //In order to create a new entry, we need the biggest Id
    //of all of the elements. Unfortunatly we cannot do that inside the
    //template of the go, so we are doing it here.
    var biggestId int
    for _, it := range t.Entry1 {
        if it.Id > biggestId {
            biggestId = it.Id
        }
    }

    biggestId = biggestId + 1;
    data := map[string]interface{}{
        "t": t,
        "akademicYearId": akademicYearId,
        "biggestId": biggestId,
    }

	RenderTemplate("teacher/report", helpers, resp, data)
}

func GetTeacherListReport(resp http.ResponseWriter, req *http.Request) {
    var helpers []string

    teacher := models.GetTeacherFromSession(req)
    akademicYears := models.ListAkademicYears()
    akademicYearsLen := len(akademicYears)
    lastFiveYears := make([]models.AkademicYear, 0)

    //TODO Wood 5 code. What we want to do here is to pass only the last five
    //years into to an array in order to print them into our template.
    //Instead of that we are iterating everything, and we are appending only
    //the last five. Patches are always welcome :)
    if akademicYearsLen > 5 {
        for index, it := range akademicYears {
            if index + 5 >= akademicYearsLen {
                lastFiveYears = append(lastFiveYears, it)
            }
        }
    } else {
        //we have less than 5 akademic reports, so show everything
        lastFiveYears = akademicYears
    }

    data := map[string]interface{}{
        "t": teacher,
        "akademicYears": akademicYears,
        "lastFiveYears": lastFiveYears,
    }

	RenderTemplate("teacher/report_list", helpers, resp, data)
}

func TeacherCreateReport1(resp http.ResponseWriter, req *http.Request) {
	teacherId := models.GetTeacherFromSession(req).Id
    akademicYearId, _ := getAkademicYearId(req)

	req.ParseForm()

	form := &forms.TeacherCreateReportFormEntry1{}
	parseTeacherCreateReportFormEntry1(form, req)

	form.Create(teacherId, akademicYearId)

    http.Redirect(resp, req, "/teacher/report/list", http.StatusMovedPermanently)
}

func TeacherCreateReport(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	formNumber, id := getId(req)
	teacherId := models.GetTeacherFromSession(req).Id
    akademicYearId,_ := getAkademicYearId(req)

	decoder := schema.NewDecoder()

	var err error
	switch formNumber {
	case 2:
		form := &forms.TeacherCreateReportFormEntry2{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId, akademicYearId)
	case 3:
		form := &forms.TeacherCreateReportFormEntry3{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId, akademicYearId)
	case 4:
		form := &forms.TeacherCreateReportFormEntry4{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId, akademicYearId)
	case 5:
		form := &forms.TeacherCreateReportFormEntry5{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(teacherId, akademicYearId)
	default:
		unknownErr := "Unknown form: " + id
		panic(unknownErr)
	}

	if err != nil {
		panic(err)
	}
    http.Redirect(resp, req, "/teacher/report/list", http.StatusMovedPermanently)
}

func GetTeacherCreateReport1Edit(resp http.ResponseWriter, req *http.Request) {
    entryId, _ := getId(req)
    akademicYearId, _ := getAkademicYearId(req)
	var helpers []string
    formsList := forms.ListAllTeacherCreateReportForm1(models.GetTeacherFromSession(req).Id, akademicYearId)

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
    akademicYearId, _ := getAkademicYearId(req)

	form.Id = id
	form.Update(akademicYearId)
    http.Redirect(resp, req, "/teacher/report/list", http.StatusMovedPermanently)
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

