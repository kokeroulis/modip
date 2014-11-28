package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
    "fmt"
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
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	form := &models.ResearchProgramCreateReportForm{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(form, req.PostForm)

	if err != nil {
		panic(err)
	}

    fmt.Println(form)
	l := models.ResearchProgramCreateReportForm{
        analutika_stoixeia_programmatos_titlos_programmatos : form.analutika_stoixeia_programmatos_titlos_programmatos,
        analutika_stoixeia_programmatos_ylopoihsh : form.analutika_stoixeia_programmatos_ylopoihsh,
        analutika_stoixeia_programmatos_foreas_xrimatodotisis : form.analutika_stoixeia_programmatos_foreas_xrimatodotisis,
        analutika_stoixeia_programmatos_proupologismos : form.analutika_stoixeia_programmatos_proupologismos,
        analutika_stoixeia_programmatos_diarkia : form.analutika_stoixeia_programmatos_diarkia,
        analutika_stoixeia_programmatos_hmerominia : form.analutika_stoixeia_programmatos_hmerominia,
        analutika_stoixeia_programmatos_sxolia : form.analutika_stoixeia_programmatos_sxolia,
	}

	l.Create()

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

