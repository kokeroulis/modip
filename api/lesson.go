package api

import (
	"github.com/kokeroulis/modip/models"
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

func GetLessonPreDegreeCreateReport (resp http.ResponseWriter, req *http.Request) {
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
        "templates/lesson/create_report/dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos.tmpl",
        "templates/lesson/create_report/aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion.tmpl",
        "templates/lesson/create_report/ekpedeutika_mesa.tmpl",
        "templates/lesson/create_report/statistoika_stoixeia_spoudaston.tmpl",
        "templates/lesson/create_report/katanomi_ba8mon_spoudaston.tmpl",
        "templates/lesson/create_report/apopsi_spoudaston_gia_to_ma8ima.tmpl",
        "templates/lesson/create_report/sxolia.tmpl",
	}

	RenderTemplate("lesson/create_report/create_report", helpers, resp, models.ListLessonsPreDegree())
}

func LessonPreDegreeCreateReport (resp http.ResponseWriter, req *http.Request) {
}

