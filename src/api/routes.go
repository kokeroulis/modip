package api

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"types"
)

type authorize struct {
	handler func(resp http.ResponseWriter, req *http.Request)
}

func (auth authorize) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	teacher := types.GetTeacherFromSession(req)
	if teacher.Id == 0 {
		Render.JSON(resp, http.StatusForbidden, map[string]string{"error": "Access denied"})
		return
	}

	auth.handler(resp, req)
}

func setupRoutes(n *negroni.Negroni) {
	router := mux.NewRouter()

	router.HandleFunc("/", Login).Methods("GET")
	router.HandleFunc("/teacher/login",TeacherLogin).Methods("POST")
	router.Handle("/paper/add", authorize{PaperAdd}).Methods("POST")

	n.UseHandler(router)
}
