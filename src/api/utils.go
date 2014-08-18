package api

import (
	"net/http"
	"database/sql"
)

func RenderJson(resp http.ResponseWriter, json interface{}) {
	Render.JSON(resp, http.StatusOK, json)
}

func checkQuery(err error) (bool) {
	if err == sql.ErrNoRows || err != nil {
		return false
	} else {
		return true
	}
}

