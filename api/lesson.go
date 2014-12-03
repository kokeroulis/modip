package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/kokeroulis/modip/models/forms"
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


func LessonListPreDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string

    data := map[string]interface{}{
        "Departments": models.ListAllDepartments(false),
        "AkademicYears": models.ListAkademicYears(),
    }

	RenderTemplate("lesson/list", helpers, resp, data)
}

func LessonListPreDegreeDepartment(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]

	departmentId, err := strconv.Atoi(id)

	if err != nil || !ok {
		return
	}

	d := models.Department{
		Id: departmentId,
	}

	d.LoadLessons(false)

    akademicYearId, _ := getAkademicYearId(req)
    data := map[string]interface{}{
        "lessons": d.Lessons,
        "akademicYearId": akademicYearId,
    }

	var helpers []string
	RenderTemplate("lesson/list_department", helpers, resp, data)
}

func LessonListPostDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/list", helpers, resp, models.ListLessonsPostDegree())
}

func GetLessonPreDegreeCreateReport (resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["lesson_id"]

	lessonId, paramErr := strconv.Atoi(id)

	if paramErr != nil || !ok {
		panic(paramErr)
	}

	l := forms.LessonCreateReportForm{
		LessonId: lessonId,
	}

	l.Load()

    akademicYearId, _ := getAkademicYearId(req)
    data := map[string]interface{}{
        "lessons": l,
        "akademicYearId": akademicYearId,
    }

	helpers := []string{
		"templates/lesson/create_report/perigrafi.tmpl",
		"templates/lesson/create_report/enhmerwsi.tmpl",
        "templates/lesson/create_report/didaktea_yli.tmpl",
        "templates/lesson/create_report/didaskalia.tmpl",
        "templates/lesson/create_report/didaktika_boh8imata.tmpl",
        "templates/lesson/create_report/epikoinonia.tmpl",
        "templates/lesson/create_report/alles_ekpedeutikes_drastiriotites.tmpl",
        "templates/lesson/create_report/summetoxi_spoudaston.tmpl",
        "templates/lesson/create_report/aksiologisi_ths_epidosis.tmpl",
        "templates/lesson/create_report/dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos.tmpl",
        "templates/lesson/create_report/aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion.tmpl",
        "templates/lesson/create_report/ekpedeutika_mesa.tmpl",
        "templates/lesson/create_report/statistoika_stoixeia_spoudaston.tmpl",
        "templates/lesson/create_report/katanomi_ba8mon_spoudaston.tmpl",
        "templates/lesson/create_report/apopsi_spoudaston_gia_to_ma8ima.tmpl",
        "templates/lesson/create_report/sxolia.tmpl",
	}

	RenderTemplate("lesson/create_report/create_report", helpers, resp, data)
}

func GetLessonEdit (resp http.ResponseWriter, req *http.Request) {
    var helpers []string

	id, _ := getId(req)

	l := models.Lesson{
		Id: id,
	}

	l.Load()
    RenderTemplate("lesson/edit", helpers, resp, l)
}

func LessonEdit (resp http.ResponseWriter, req *http.Request) {
	l := &models.Lesson{}

	req.ParseForm()
	decoder := schema.NewDecoder()

	err := decoder.Decode(l, req.PostForm)

	if err != nil {
		panic(err)
	}

	l.Update()

	url := "lesson/edit/" + string(l.Id)
	http.Redirect(resp, req, url, http.StatusMovedPermanently)
}

func LessonPreDegreeCreateReport (resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	id2, ok2 := vars["lesson_id"]
	id3, ok3 := vars["akademicYearId"]

	formNumber, paramErr := strconv.Atoi(id)
	lessonId, paramErr2 := strconv.Atoi(id2)
    _, paramErr3 := strconv.Atoi(id3)

	if paramErr != nil || !ok {
		panic(paramErr)
	}

	if !ok2 || paramErr2 != nil {
		panic(paramErr2)
	}

	if !ok3 || paramErr3 != nil {
		panic(paramErr3)
	}

	decoder := schema.NewDecoder()
	req.ParseForm()

	var err error

	switch formNumber {
	case 1:
		form := &forms.LessonCreateReportFormEntry1{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 2:
		form := &forms.LessonCreateReportFormEntry2{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 3:
		form := &forms.LessonCreateReportFormEntry3{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 4:
		form := &forms.LessonCreateReportFormEntry4{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 5:
		form := &forms.LessonCreateReportFormEntry5{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 11:
		form := &forms.LessonCreateReportFormEntry11{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 12:
		form := &forms.LessonCreateReportFormEntry12{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 13:
		form := &forms.LessonCreateReportFormEntry13{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 14:
		form := &forms.LessonCreateReportFormEntry14{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 15:
		form := &forms.LessonCreateReportFormEntry15{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	case 16:
		form := &forms.LessonCreateReportFormEntry16{}
		err = decoder.Decode(form, req.PostForm)
		form.Update(lessonId)
	default:
		unknownErr := "Unknown form: " + id
		panic(unknownErr)
	}

	if err != nil {
		panic(err)
	}

	url := "/lesson/pre/degree/create/report/" + id2 + "/" + id3
	http.Redirect(resp, req, url , http.StatusMovedPermanently)
}

