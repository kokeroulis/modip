package api

import (
	"github.com/kokeroulis/modip/models"
	"net/http"
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

