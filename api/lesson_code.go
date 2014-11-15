package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
)

type LessonCodeCreateForm struct {
	LessonName    string `schema:"lesson_code_create_lesson_name"`
	Department    int    `schema:"lesson_code_create_lesson_department"`
	CourseCode    string `schema:"lesson_code_create_lesson_code"`
	CardisoftCode string `schema:"lesson_code_create_lesson_cardisoft_code"`
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

	l := models.Lesson{
		Name:          form.LessonName,
		Department:    form.Department,
		CourseCode:    form.CourseCode,
		CardisoftCode: form.CardisoftCode,
		IsPostDegree:  false,
	}

	l.Create()
	http.Redirect(resp, req, "/lesson/code/list", http.StatusMovedPermanently)
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

