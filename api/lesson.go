package api
import (
	"github.com/kokeroulis/modip/models"
	"github.com/mholt/binding"
	"net/http"
	"fmt"
)


func LessonListPreDegree(resp http.ResponseWriter, req *http.Request) {
	RenderTemplate("lesson/list", resp, models.ListLessonsPreDegree())
}

func LessonListPostDegree(resp http.ResponseWriter, req *http.Request) {
	RenderTemplate("lesson/list", resp, models.ListLessonsPostDegree())
}

type LessonSaveForm struct {
	LessonName   string
	Department   int
	IsPostDegree bool
}

func (l *LessonSaveForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&l.LessonName: binding.Field{
			Form:     "lesson_name",
			Required: true,
		},
		&l.Department: binding.Field{
			Form:     "department",
			Required: true,
		},
		&l.IsPostDegree: binding.Field{
			Form:     "is_post_degree",
			Required: true,
		},
	}
}

func LessonCreate(resp http.ResponseWriter, req *http.Request) {
	lessonSaveForm := &LessonSaveForm{}
	fmt.Println("here")
	if binding.Bind(req, lessonSaveForm).Handle(resp) {
		return
	}

	l := models.Lesson{
		Name: lessonSaveForm.LessonName,
		Department: lessonSaveForm.Department,
		IsPostDegree: lessonSaveForm.IsPostDegree,
	}

	fmt.Println(l)

	http.Redirect(resp, req, "/", 200)
}
