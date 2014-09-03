package api

import (
	"net/http"
	"database/sql"
	"fmt"
	"types"
)

func RenderJson(resp http.ResponseWriter, json interface{}) {
	Render.JSON(resp, http.StatusOK, json)
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
	} else {
		noRows = false
		dbError = false
	}

	return dbError, noRows
}

func renderDbError(resp http.ResponseWriter, req *http.Request) {
	errorJson := types.DbError()
	dbErrorJson := types.CreateStandardJsonErrorJson(req, errorJson)
	Render.JSON(resp, errorJson.Code, dbErrorJson)
}

