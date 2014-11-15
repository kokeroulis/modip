package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


func LessonListPreDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/list", helpers, resp, models.ListAllDepartments(false))
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

	var helpers []string
	RenderTemplate("lesson/list_department", helpers, resp, d.Lessons)
}

func LessonListPostDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/list", helpers, resp, models.ListLessonsPostDegree())
}

type LessonCreateForm struct {
	LessonName   string `schema:"lesson_name"`
	Department   int    `schema:"department"`
	IsPostDegree bool   `schema:"is_post_degree"`
}

func LessonCreate(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	form := &LessonCreateForm{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(form, req.PostForm)

	if err != nil {
		panic(err)
	}

}

func GetLessonPreDegreeCreate (resp http.ResponseWriter, req *http.Request) {
	helpers := []string{
		"templates/lesson/create_report/perigrafi.tmpl",
		"templates/lesson/create_report/didaskalia.tmpl",
		"templates/lesson/create_report/enhmerwsi.tmpl",
        "templates/lesson/create_report/didaktea_yli.tmpl",
        "templates/lesson/create_report/didaktika_boh8imata.tmpl",
        "templates/lesson/create_report/epikoinonia.tmpl",
        "templates/lesson/create_report/alles_ekpedeutikes_drastiriotites.tmpl",
        "templates/lesson/create_report/summetoxi_spoudaston.tmpl",
        "templates/lesson/create_report/aksiologisi_ths_epidosis.tmpl",
	}

	RenderTemplate("lesson/create_report/create_report", helpers, resp, models.ListLessonsPreDegree())
}

