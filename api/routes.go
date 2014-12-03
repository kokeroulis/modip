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
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")

	router.Handle("/category/list", authorize{CategoryList}).Methods("GET")
	router.Handle("/category/list/{id:[1-9]+}", authorize{CategoryList}).Methods("GET")

	router.Handle("/category/save", authorize{CategorySave}).Methods("POST")

    router.Handle("/akademic/year/list", authorize{GetAkademicYearList}).Methods("GET")
    router.Handle("/akademic/year/create", authorize{GetAkademicYearCreate}).Methods("GET")
    router.Handle("/akademic/year/create", authorize{AkademicYearCreate}).Methods("POST")
    router.Handle("/akademic/year/edit/{id:[1-9]+}", authorize{GetAkademicYearEdit}).Methods("GET")
    router.Handle("/akademic/year/edit", authorize{AkademicYearEdit}).Methods("POST")

	// Α.Δ. Εκπ. Προσωπικού
	router.Handle("/teacher/report/list", authorize{GetTeacherListReport}).Methods("GET")
	router.Handle("/teacher/report/{akademicYearId:[0-9]+}", authorize{GetTeacherCreateReport}).Methods("GET")
	router.Handle("/teacher/report/1", authorize{TeacherCreateReport1}).Methods("POST")
	router.Handle("/teacher/report/1/edit", authorize{GetTeacherCreateReport1Edit}).Methods("GET")
	router.Handle("/teacher/report/1/edit/{id:[1-9]+}", authorize{TeacherCreateReport1Edit}).Methods("POST")
        router.Handle("/teacher/report/1/edit/{id:[0-9],akademicYearId:[0-9]+}", authorize{GetTeacherCreateReport1Edit}).Methods("GET")
        router.Handle("/teacher/report/1/edit/{id:[0-9],akademicYearId:[0-9]+}", authorize{TeacherCreateReport1Edit}).Methods("POST")

	router.Handle("/teacher/report/{id:[2-5]}/{akademicYearId:[0-9]+}", authorize{TeacherCreateReport}).Methods("POST")

	// Α.Δ. Ερευν. Προγραμ.
	router.Handle("/research/program", authorize{GetResearchProgram}).Methods("GET")
	router.Handle("/research/program/create/report", authorize{GetResearchProgramCreateReport}).Methods("GET")
	router.Handle("/research/program/create/report", authorize{ResearchProgramCreateReport}).Methods("POST")

	router.Handle("/research/program/edit/report/{id:[1-9]+}", authorize{GetResearchProgramEditReport}).Methods("GET")
	router.Handle("/research/program/edit/report/{id:[1-9]+}", authorize{ResearchProgramEditReport}).Methods("POST")

	// Α.Δ. Μαθημάτων Π.Π.Σ.
	router.Handle("/lesson/list/pre/degree", authorize{LessonListPreDegree}).Methods("GET")
	router.Handle("/lesson/list/pre/degree/{id:[1-9]+}", authorize{LessonListPreDegreeDepartment}).Methods("GET")
	router.Handle("/lesson/pre/degree/create/report/{lesson_id:[1-9]+}", authorize{GetLessonPreDegreeCreateReport}).Methods("GET")
	router.Handle("/lesson/pre/degree/create/report/{id:[0-9]+}/{lesson_id:[0-9]+}", authorize{LessonPreDegreeCreateReport}).Methods("POST")

	router.Handle("/lesson/list/post/degree", authorize{LessonListPostDegree}).Methods("GET")


	// Κωδ. Μαθημάτων Π.Π.Σ.
	router.Handle("/lesson/code/list", authorize{GetLessonCodeList}).Methods("GET")
	router.Handle("/lesson/code/list/{id:[0-9]+}", authorize{GetLessonCodeListDepartment}).Methods("GET")
	router.Handle("/lesson/code/create", authorize{GetLessonCodeCreate}).Methods("GET")
	router.Handle("/lesson/code/create", authorize{LessonCodeCreate}).Methods("POST")
	router.Handle("/lesson/edit/{id:[0-9]+}", authorize{GetLessonEdit}).Methods("GET")
	router.Handle("/lesson/edit", authorize{LessonEdit}).Methods("POST")

    //Προσωπικό Μονάδας
    router.Handle("/stuff/list", authorize{GetStuffList}).Methods("GET")
    router.Handle("/stuff/create", authorize{GetStuffCreate}).Methods("GET")
    router.Handle("/stuff/create", authorize{StuffCreate}).Methods("POST")
    router.Handle("/stuff/edit/{id:[1-9]+}", authorize{GetStuffEdit}).Methods("GET")
    router.Handle("/stuff/edit", authorize{StuffEdit}).Methods("POST")

	n.UseHandler(router)
}
