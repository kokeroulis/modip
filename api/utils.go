package api

import (
	"github.com/kokeroulis/modip/models"
	"net/http"
	"html/template"
	"strings"
)

func RenderJson(resp http.ResponseWriter, req *http.Request, j interface{}) {
	data := models.JsonData{
		Common: models.CreateStandardJson(req),
		Data:   j,
	}

	Render.JSON(resp, http.StatusOK, data)
}

func RenderErrorJson(resp http.ResponseWriter, req *http.Request, errorJson models.ErrorJson, j interface{}) {
	data := models.JsonData{
		Common: models.CreateStandardJsonErrorJson(req, errorJson),
		Data:   j,
	}

	Render.JSON(resp, errorJson.Code, data)
}

func RenderTemplate(name string, helpers []string, resp http.ResponseWriter, data interface{}) {
	templateFile := "templates/" + name + ".tmpl"

	var templates []string
	templates = append(templates, templateFile)

	templates = append(templates, "templates/base/header.tmpl")
	templates = append(templates, "templates/base/footer.tmpl")
	templates = append(templates, "templates/base/sidebar.tmpl")


	tmpl, err := template.New("").ParseFiles(templates...)

	if err != nil {
		panic(err)
	}

	tmplListPath := strings.Split(name, "/")
	err = tmpl.ExecuteTemplate(resp, tmplListPath[len(tmplListPath) - 1] + ".tmpl", data)

	if err != nil {
		panic(err)
	}
}

