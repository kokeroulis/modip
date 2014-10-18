package models

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/boj/redistore"
	"github.com/kokeroulis/modip/config"
	"net/http"
)

var redisStore *redistore.RediStore

type AuthJson struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type CommonJson struct {
	Auth  AuthJson  `json:"auth"`
	Error ErrorJson `json:"error"`
}

type JsonData struct {
	Common CommonJson
	Data   interface{} `json:"data"`
}

func SetupRedis() {
	c := config.NewConfig()

	var err error
	redisStore, err = redistore.NewRediStore(10, "tcp", c.RedisPort, "", c.RedisSecret)

	if err != nil {
		fmt.Println("Can't connect to redis")
		panic(err)
	}
}

func GetSession(req *http.Request) *sessions.Session {
	session, err := redisStore.Get(req, "my_cookie")

	if err != nil {
		fmt.Println("Error at getting session!!!")
		fmt.Println(err)
	}

	return session
}

func GetTeacherFromSession(req *http.Request) Teacher {
	session := GetSession(req)

	t := session.Values["teacher"]
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

type BookOrPaperInfo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
