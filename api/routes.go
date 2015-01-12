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
    teacherId := teacher.Id
    teacherType := teacher.Type
	if teacherId == 0 {
		panic("TODO!!!!!!!!!!!!!!1")
		return
    } else if teacherType == 12 {
        auth.handler(resp, req)
        return
    } else if teacherType != 1 && teacherType != 2 && teacherType != 3 && teacherType != 4 && teacherType != 5 && teacherType != 6 && teacherType != 10 {
        http.Redirect(resp, req, "/lesson/list/pre/degree", http.StatusMovedPermanently)
        return
    }

	auth.handler(resp, req)
}

type authorizeSecretary struct {
	handler func(resp http.ResponseWriter, req *http.Request)
}

func (auth authorizeSecretary) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	teacher := models.GetTeacherFromSession(req)
    teacherType := teacher.Type
	if teacher.Id == 0 {
		panic("TODO!!!!!!!!!!!!!!1")
		return
    } else if teacherType == 12 {
        auth.handler(resp, req)
        return
    } else if teacherType != 7 && teacherType != 8 && teacherType != 9 && teacherType != 11 {
        http.Redirect(resp, req, "/teacher/report/list", http.StatusMovedPermanently)
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

    router.Handle("/akademic/year/list", authorizeSecretary{GetAkademicYearList}).Methods("GET")
    router.Handle("/akademic/year/create", authorizeSecretary{GetAkademicYearCreate}).Methods("GET")
    router.Handle("/akademic/year/create", authorizeSecretary{AkademicYearCreate}).Methods("POST")
    router.Handle("/akademic/year/edit/{id:[1-9]+}", authorizeSecretary{GetAkademicYearEdit}).Methods("GET")
    router.Handle("/akademic/year/edit", authorizeSecretary{AkademicYearEdit}).Methods("POST")

	// Α.Δ. Εκπ. Προσωπικού
	router.Handle("/teacher/report/list", authorize{GetTeacherListReport}).Methods("GET")
	router.Handle("/teacher/report/{akademicYearId:[0-9]+}", authorize{GetTeacherCreateReport}).Methods("GET")
	router.Handle("/teacher/report/1/{akademicYearId:[0-9]+}", authorize{TeacherCreateReport1}).Methods("POST")
    router.Handle("/teacher/report/1/edit/{id:[0-9]}/{akademicYearId:[0-9]+}", authorize{GetTeacherCreateReport1Edit}).Methods("GET")
	router.Handle("/teacher/report/1/edit/{id:[0-9]}/{akademicYearId:[0-9]+}", authorize{TeacherCreateReport1Edit}).Methods("POST")

	router.Handle("/teacher/report/{id:[2-5]}/{akademicYearId:[0-9]+}", authorize{TeacherCreateReport}).Methods("POST")

	// Α.Δ. Ερευν. Προγραμ.
	router.Handle("/research/program", authorize{GetResearchProgram}).Methods("GET")
	router.Handle("/research/program/create/report", authorize{GetResearchProgramCreateReport}).Methods("GET")
	router.Handle("/research/program/create/report", authorize{ResearchProgramCreateReport}).Methods("POST")

	router.Handle("/research/program/edit/report/{id:[1-9]+}", authorize{GetResearchProgramEditReport}).Methods("GET")
	router.Handle("/research/program/edit/report/{id:[1-9]+}", authorize{ResearchProgramEditReport}).Methods("POST")

	// Α.Δ. Μαθημάτων Π.Π.Σ.
	router.Handle("/lesson/list/pre/degree", authorizeSecretary{LessonListPreDegree}).Methods("GET")
	router.Handle("/lesson/list/pre/degree/{id:[1-9]+}/{akademicYearId:[1-9]+}", authorizeSecretary{LessonListPreDegreeDepartment}).Methods("GET")
	router.Handle("/lesson/pre/degree/create/report/{lesson_id:[0-9]+}/{akademicYearId:[1-9]+}", authorizeSecretary{GetLessonPreDegreeCreateReport}).Methods("GET")
	router.Handle("/lesson/pre/degree/create/report/{id:[0-9]+}/{lesson_id:[0-9]+}/{akademicYearId:[1-9]+}", authorizeSecretary{LessonPreDegreeCreateReport}).Methods("POST")

	router.Handle("/lesson/list/post/degree", authorizeSecretary{LessonListPostDegree}).Methods("GET")


	// Κωδ. Μαθημάτων Π.Π.Σ.
	router.Handle("/lesson/code/list", authorizeSecretary{GetLessonCodeList}).Methods("GET")
	router.Handle("/lesson/code/list/{id:[0-9]+}", authorizeSecretary{GetLessonCodeListDepartment}).Methods("GET")
	router.Handle("/lesson/code/create", authorizeSecretary{GetLessonCodeCreate}).Methods("GET")
	router.Handle("/lesson/code/create", authorizeSecretary{LessonCodeCreate}).Methods("POST")
	router.Handle("/lesson/edit/{id:[0-9]+}", authorizeSecretary{GetLessonEdit}).Methods("GET")
	router.Handle("/lesson/edit", authorizeSecretary{LessonEdit}).Methods("POST")

    //Προσωπικό Μονάδας
    router.Handle("/stuff/list", authorizeSecretary{GetStuffList}).Methods("GET")
    router.Handle("/stuff/create", authorizeSecretary{GetStuffCreate}).Methods("GET")
    router.Handle("/stuff/create", authorizeSecretary{StuffCreate}).Methods("POST")
    router.Handle("/stuff/edit/{id:[1-9]+}", authorizeSecretary{GetStuffEdit}).Methods("GET")
    router.Handle("/stuff/edit", authorizeSecretary{StuffEdit}).Methods("POST")

	n.UseHandler(router)
}
