package models

import "net/http"

type ErrorJson struct {
	Code int
	Name string
}

func AuthFailed() ErrorJson {
	return ErrorJson{http.StatusUnauthorized, "Authorization Failed"}
}

func AlreadyExists() ErrorJson {
	return ErrorJson{http.StatusOK, "AlreadyExists"}
}

func DbError() ErrorJson {
	return ErrorJson{http.StatusOK, "DbError"}
}

func InvalidAsset() ErrorJson {
	return ErrorJson{http.StatusOK, "InvalidAsset"}
}

