package api

import (
	"net/http"
	"database/sql"
	"fmt"
)

func RenderJson(resp http.ResponseWriter, json interface{}) {
	Render.JSON(resp, http.StatusOK, json)
}

func checkQuery(err error) (bool) {
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

