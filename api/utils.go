package api

import (
	"database/sql"
	"fmt"
	"github.com/kokeroulis/modip/models"
	"net/http"
)

func RenderJson(resp http.ResponseWriter, json interface{}) {
	Render.JSON(resp, http.StatusOK, json)
}

func RenderJson2(resp http.ResponseWriter, req *http.Request, j interface{}) {
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

// we don't really use resp or req but they are usefull for
// debugging purposes.
// So we will leave them as parameters
func checkQuery(err error, resp http.ResponseWriter,
	req *http.Request, query string) (bool, bool) {
	var noRows bool
	var dbError bool

	if err == sql.ErrNoRows {
		dbError = false
	} else if err != nil {
		// print the error
		// TODO use logger
		fmt.Println(err)
		dbError = true
		noRows = true
		renderDbError(resp, req)
	} else {
		noRows = false
		dbError = false
	}

	return dbError, noRows
}

func renderDbError(resp http.ResponseWriter, req *http.Request) {
	errorJson := models.DbError()
	dbErrorJson := models.CreateStandardJsonErrorJson(req, errorJson)
	Render.JSON(resp, errorJson.Code, dbErrorJson)
}
