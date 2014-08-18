package types

import "net/http"

type ErrorJson struct {
	Code int
	Name string
}

func AuthFailed() ErrorJson {
	return ErrorJson{http.StatusUnauthorized, "Authorization Failed"}
}

