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
		panic("TODO!!!!!!!!!!!!!!1")
		return
	}

	auth.handler(resp, req)
}

func setupRoutes(n *negroni.Negroni) {
	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")

	router.HandleFunc("/teacher/login", TeacherLogin).Methods("POST")

	router.Handle("/category/list", authorize{CategoryList}).Methods("GET")
	router.Handle("/category/list/{id:[1-9]+}", authorize{CategoryList}).Methods("GET")

	router.Handle("/category/save", authorize{CategorySave}).Methods("POST")

	// Α.Δ. Εκπ. Προσωπικού
	router.Handle("/teacher/report", authorize{GetTeacherCreateReport}).Methods("GET")
	router.Handle("/teacher/report/1", authorize{TeacherCreateReport1}).Methods("POST")
	router.Handle("/teacher/report/{id:[2-5]}", authorize{TeacherCreateReport}).Methods("POST")

	// Α.Δ. Μαθημάτων Π.Π.Σ.
	router.Handle("/lesson/list/pre/degree", authorize{LessonListPreDegree}).Methods("GET")
	router.Handle("/lesson/list/pre/degree/{id:[1-9]+}", authorize{LessonListPreDegreeDepartment}).Methods("GET")
	router.Handle("/lesson/pre/degree/create/report/{lesson_id:[0-9]+}", authorize{GetLessonPreDegreeCreateReport}).Methods("GET")
	router.Handle("/lesson/pre/degree/create/report/{id:[0-9]+}/{lesson_id:[0-9]+}", authorize{LessonPreDegreeCreateReport}).Methods("POST")

	router.Handle("/lesson/list/post/degree", authorize{LessonListPostDegree}).Methods("GET")


	// Κωδ. Μαθημάτων Π.Π.Σ.
	router.Handle("/lesson/code/list", authorize{LessonCodeList}).Methods("GET")
	router.Handle("/lesson/code/create", authorize{GetLessonCodeCreate}).Methods("GET")
	router.Handle("/lesson/code/create", authorize{LessonCodeCreate}).Methods("POST")

	n.UseHandler(router)
}
