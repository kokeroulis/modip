package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
)

type LessonCodeCreateForm struct {
	LessonName   string `schema:"lesson_name"`
	Department   int    `schema:"department"`
	IsPostDegree bool   `schema:"is_post_degree"`
	CourseCode   string   `schema:course_code`
}

func LessonCodeCreate(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	form := &LessonCodeCreateForm{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(form, req.PostForm)

	if err != nil {
		panic(err)
	}

}

func GetLessonCodeCreate(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/code/create", helpers, resp, models.ListAllDepartments(false))
}

func LessonCodeList(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}
	RenderTemplate("lesson/code/list", helpers, resp, data)
}

