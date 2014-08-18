package types

import (
	"net/http"
	"github.com/goincremental/negroni-sessions"
	"fmt"
)

type AuthJson struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type CommonJson struct {
	Auth  AuthJson  `json:"auth"`
	Error ErrorJson `json:"error"`
}

func GetTeacherFromSession(req *http.Request) Teacher {
	t := sessions.GetSession(req).Get("teacher")
	teacher := &Teacher{}
	if t != nil {
		ok := false
		teacher, ok = t.(*Teacher)
		if !ok {
			fmt.Println("Type assertion failed")
			fmt.Println("Unauthorized request")
		}
	}

	return *teacher
}

func createAuth(req *http.Request) AuthJson {
	teacher := GetTeacherFromSession(req)

	auth := AuthJson{teacher.Id, teacher.Email}

	return auth
}

func CreateStandardJson(req *http.Request) CommonJson {
	return CommonJson{createAuth(req), ErrorJson{}}
}

func CreateStandardJsonErrorJson(req *http.Request, errorJson ErrorJson) CommonJson {
	return CommonJson{createAuth(req), errorJson}
}

