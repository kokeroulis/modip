package api
import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
)


func LessonListPreDegree(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	RenderTemplate("lesson/list", helpers, resp, models.ListLessonsPreDegree())
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
	var helpers []string
	RenderTemplate("lesson/create/create", helpers, resp, models.ListLessonsPreDegree())
}

