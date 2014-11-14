package api
import (
//	"encoding/gob"
	"github.com/kokeroulis/modip/models"
//	"github.com/mholt/binding"
	"net/http"
)


func LessonListPreDegree(resp http.ResponseWriter, req *http.Request) {
	RenderTemplate("lesson/list", resp, models.ListLessonsPreDegree())
}

func LessonListPostDegree(resp http.ResponseWriter, req *http.Request) {
	RenderTemplate("lesson/list", resp, models.ListLessonsPostDegree())
}

