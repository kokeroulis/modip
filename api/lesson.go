package api
import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
	"fmt"
)


func LessonListPreDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/list", helpers, resp, models.ListLessonsPreDegree())
}

func LessonListPostDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/list", helpers, resp, models.ListLessonsPostDegree())
}

type LessonSaveForm struct {
	LessonName   string `schema:"lesson_name"`
	Department   int    `schema:"department"`
	IsPostDegree bool   `schema:"is_post_degree"`
}

func LessonCreate(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	lessonSaveForm := &LessonSaveForm{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(lessonSaveForm, req.PostForm)

	if err != nil {
		panic(err)
	}

	fmt.Println(lessonSaveForm)
}

func GetLessonPreDegreeCreate (resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/create/create", helpers, resp, models.ListLessonsPreDegree())
}

