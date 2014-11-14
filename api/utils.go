package api

import (
	"github.com/kokeroulis/modip/models"
	"net/http"
	"io/ioutil"
	"html/template"
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

func RenderTemplate(name string, resp http.ResponseWriter, data interface{}) {
	templateName := "templates/" + name + ".tmpl"
	contents, err := ioutil.ReadFile(templateName)

	if err != nil {
		panic("Can't find template")
	}

	tmpl, err := template.New("foo").Parse(string(contents))

	if err != nil {
		panic("Can't parse the template")
	}

	err = tmpl.Execute(resp, data)

	if err != nil {
		panic("Something is wrong with executing the template!")
	}

}

