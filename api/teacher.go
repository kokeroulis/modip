package api

import (
	"encoding/gob"
	"fmt"
	"github.com/goincremental/negroni-sessions"
	"github.com/kokeroulis/modip/models"
	"github.com/mholt/binding"
	"net/http"
)

type TeacherForm struct {
	Username string
	Password string
}

func (t *TeacherForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.Username: binding.Field{
			Form:     "username",
			Required: true,
		},
		&t.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

func init() {
	// later we will try to save to our
	// session a struct. Structs are complex
	// models which aren't supported by default
	// from the god package so we have to
	// register it manually
	gob.Register(&models.Teacher{})
}

func TeacherLogin(resp http.ResponseWriter, req *http.Request) {
	teacherForm := &TeacherForm{}
	if binding.Bind(req, teacherForm).Handle(resp) {
		return
	}

	t := models.Teacher{}
	query := `SELECT id, name, email, departmentId, departmentName, authFailed
			  FROM teacher_auth($1, $2)`

	var authFailed bool

	err := Db.QueryRow(query, teacherForm.Username, teacherForm.Password).
		Scan(&t.Id, &t.Name, &t.Email, &t.Department.Id, &t.Department.Name, &authFailed)

	var dbError, noRows = checkQuery(err, resp, req, query)

	if dbError {
		return
	}

	if !noRows && !authFailed {
		session := sessions.GetSession(req)
		session.Set("teacher", &t)
		RenderJson2(resp, req, t)
	} else {
		errorJson := models.AuthFailed()
		RenderErrorJson(resp, req, errorJson, t)
	}
}

func retrieveInfo(teacherId int, query string, channel chan []models.BookOrPaperInfo) {
	list := []models.BookOrPaperInfo{}
	rows, err := Db.Query(query, teacherId)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := models.BookOrPaperInfo{}

		if err := rows.Scan(&it.Id, &it.Title); err != nil {
			fmt.Println(err)
		} else {
			list = append(list, it)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	channel <- list
}

func TeacherInfo(resp http.ResponseWriter, req *http.Request) {
	teacher := models.GetTeacherFromSession(req)

	books := make(chan []models.BookOrPaperInfo)
	booksQuery := `SELECT id, title FROM book WHERE teacher = $1`

	papers := make(chan []models.BookOrPaperInfo)
	papersQuery := `SELECT id, title FROM paper WHERE teacher = $1`

	go retrieveInfo(teacher.Id, booksQuery, books)
	go retrieveInfo(teacher.Id, papersQuery, papers)

	info := models.TeacherInfo{teacher, <-books, <-papers}

	RenderJson2(resp, req, info)
}
