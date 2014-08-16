package api

import (
	"net/http"
	"database/sql"
)

func RenderJson(resp http.ResponseWriter, json interface{}) {
	Render.JSON(resp, http.StatusOK, json)
}

type ErrorJson struct {
	Code int
	Name string
}

func authFailed() ErrorJson {
	return ErrorJson{http.StatusUnauthorized, "Authorization Failed"}
}

type myFunc func() ErrorJson

func RenderError(jsonError myFunc, resp http.ResponseWriter) {
	err := jsonError()
	Render.JSON(resp, err.Code, map[string]string{"error": err.Name})
}

func checkQuery(err error) (bool) {
	if err == sql.ErrNoRows || err != nil {
		return false
	} else {
		return true
	}
}

