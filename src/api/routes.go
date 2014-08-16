package api

import (
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler func(resp http.ResponseWriter, req *http.Request)

type Authorize struct {
	handler func(resp http.ResponseWriter, req *http.Request)
}

func (auth *Authorize) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	session := sessions.GetSession(req)
	if session.Get("foo") == nil {
		Render.JSON(resp, http.StatusForbidden, map[string]string{"error": "Access denied"})
		return
	}

	auth.handler(resp, req)
}

func setupRoutes(n *negroni.Negroni) {
	router := mux.NewRouter()

	router.HandleFunc("/", Login).Methods("GET")
	router.HandleFunc("/teacher/login",TeacherLogin).Methods("POST")

	n.UseHandler(router)
}

