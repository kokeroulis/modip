package api

import (
	_"github.com/kokeroulis/modip/models"
	_"github.com/kokeroulis/modip/models/forms"
	_"github.com/gorilla/schema"
	"net/http"
)

func GetResearchProgram(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}

	RenderTemplate("research_program/list", helpers, resp, data)
}

func GetResearchProgramCreateReport(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}

	RenderTemplate("research_program/create_report", helpers, resp, data)
}

func ResearchProgramCreateReport(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/research/program", http.StatusMovedPermanently)
}

func GetResearchProgramEditReport(resp http.ResponseWriter, req *http.Request) {
	var helpers []string
	var data interface{}

	RenderTemplate("research_program/edit_report", helpers, resp, data)
}

func ResearchProgramEditReport(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/research/program", http.StatusMovedPermanently)
}

