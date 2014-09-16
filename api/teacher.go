package api

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/mholt/binding"
	"net/http"
	"encoding/gob"
	"github.com/kokeroulis/modip/types"
	"fmt"
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
	// types which aren't supported by default
	// from the god package so we have to
	// register it manually
	gob.Register(&types.Teacher{})
}

func TeacherLogin(resp http.ResponseWriter, req *http.Request) {
	teacherForm := &TeacherForm{}
    if binding.Bind(req, teacherForm).Handle(resp) {
        return
    }

	t := types.Teacher{}
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
		teacherJson := types.TeacherJson{types.CreateStandardJson(req), t}
		RenderJson(resp, teacherJson)
	} else {
		errorJson := types.AuthFailed()
		teacherJson := types.TeacherJson{types.CreateStandardJsonErrorJson(req, errorJson), t}
		Render.JSON(resp, errorJson.Code, teacherJson)
	}
}

func retrieveInfo(teacherId int, query string, channel chan []types.BookOrPaperInfo) {
	list := []types.BookOrPaperInfo{}
	rows, err := Db.Query(query, teacherId)

   if err != nil {
		fmt.Println(err)
    }

    defer rows.Close()

	for rows.Next() {
		it := types.BookOrPaperInfo{}

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
	teacher := types.GetTeacherFromSession(req)

	books := make(chan []types.BookOrPaperInfo)
	booksQuery := `SELECT id, title FROM book WHERE teacher = $1`

	papers := make(chan []types.BookOrPaperInfo)
	papersQuery := `SELECT id, title FROM paper WHERE teacher = $1`

	go retrieveInfo(teacher.Id, booksQuery, books)
	go retrieveInfo(teacher.Id, papersQuery, papers)

	info := types.TeacherInfo{teacher, <-books, <-papers}
	infoJson := types.TeacherInfoJson{types.CreateStandardJson(req), info}

	RenderJson(resp, infoJson)
}

