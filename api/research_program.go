package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
)

func GetResearchProgram(resp http.ResponseWriter, req *http.Request) {
	var helpers []string

    teacherId :=  models.GetTeacherFromSession(req).Id
    data := models.ListAllResearchPrograms(teacherId)
	RenderTemplate("research_program/list", helpers, resp, data)
}

func GetResearchProgramCreateReport(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}

	RenderTemplate("research_program/create_report", helpers, resp, data)
}

func ResearchProgramCreateReport(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	r := &models.ResearchProgramCreateReportForm{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(r, req.PostForm)

	if err != nil {
		panic(err)
	}

	r.Create(models.GetTeacherFromSession(req).Id)

	http.Redirect(resp, req, "/research/program", http.StatusMovedPermanently)
}

func GetResearchProgramEditReport(resp http.ResponseWriter, req *http.Request) {
	var helpers []string

    r := &models.ResearchProgramCreateReportForm{}
    researchProgramId, _ := getId(req)
    r.Load(researchProgramId)

	RenderTemplate("research_program/edit_report", helpers, resp, r)
}

func ResearchProgramEditReport(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	r := &models.ResearchProgramCreateReportForm{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(r, req.PostForm)

	if err != nil {
		panic(err)
	}

    researchProgramId, _ := getId(req)
    r.Update(researchProgramId)

	http.Redirect(resp, req, "/research/program", http.StatusMovedPermanently)
}

