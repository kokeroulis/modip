package api

import (
	"github.com/kokeroulis/modip/models"
	"encoding/json"
	"net/http"
	"html/template"
	"strings"
)

var globalFunc template.FuncMap

func init() {
	globalFunc = template.FuncMap{
		"add": add,
	}
}

func add(x int, y int) int {
	return x + y
}

func RenderJson(resp http.ResponseWriter, req *http.Request, j interface{}) {
	data := models.JsonData{
		Common: models.CreateStandardJson(req),
		Data:   j,
	}

	out, _ := json.Marshal(data)
	resp.Write(out)
}

func RenderErrorJson(resp http.ResponseWriter, req *http.Request, errorJson models.ErrorJson, j interface{}) {
	data := models.JsonData{
		Common: models.CreateStandardJsonErrorJson(req, errorJson),
		Data:   j,
	}

	out, _ := json.Marshal(data)
	resp.Write(out)
}

func RenderTemplate(name string, helpers []string, resp http.ResponseWriter, data interface{}) {

	tmplListPath := strings.Split(name, "/")
	mainTemplateName := tmplListPath[len(tmplListPath) - 1] + ".tmpl"
	templateFile := "templates/" + name + ".tmpl"

	var templates []string
	templates = append(templates, templateFile)

	templates = append(templates, "templates/base/header.tmpl")
	templates = append(templates, "templates/base/footer.tmpl")
	templates = append(templates, "templates/base/sidebar.tmpl")
	templates = append(templates, helpers...)

	tmpl := template.Must(template.New(mainTemplateName).Funcs(globalFunc).ParseFiles(templates...))

	err := tmpl.ExecuteTemplate(resp, mainTemplateName, data)

	if err != nil {
		panic(err)
	}
}

