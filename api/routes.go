package api

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/kokeroulis/modip/models"
	"net/http"
)

type authorize struct {
	handler func(resp http.ResponseWriter, req *http.Request)
}

func (auth authorize) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	teacher := models.GetTeacherFromSession(req)
	if teacher.Id == 0 {
		Render.JSON(resp, http.StatusForbidden, map[string]string{"error": "Access denied"})
		return
	}

	auth.handler(resp, req)
}

func setupRoutes(n *negroni.Negroni) {
	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/test", Test).Methods("GET")

	router.HandleFunc("/teacher/login", TeacherLogin).Methods("POST")
	router.Handle("/teacher/info", authorize{TeacherInfo}).Methods("GET")

	router.Handle("/teacher/asset/add", authorize{AssetAdd}).Methods("POST")
	router.Handle("/teacher/asset/remove", authorize{AssetAdd}).Methods("POST")
	router.Handle("/teacher/asset/move", authorize{AssetAdd}).Methods("POST")
	router.Handle("/teacher/asset/modify", authorize{AssetAdd}).Methods("POST")

	n.UseHandler(router)
}
