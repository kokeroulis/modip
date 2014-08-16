package api

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/mholt/binding"
	"net/http"
	"encoding/gob"
	"types"
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
	gob.Register(&types.TeacherJson{})
}

func TeacherLogin(resp http.ResponseWriter, req *http.Request) {
	teacherForm := &TeacherForm{}
    if binding.Bind(req, teacherForm).Handle(resp) {
        return
    }

	j := types.TeacherJson{}
	query := `SELECT id, name, email, departmentId, departmentName
			  FROM teacher_auth($1, $2)`

	err := Db.QueryRow(query, teacherForm.Username, teacherForm.Password).Scan(&j.Id,
																			   &j.Name,
																			   &j.Email,
																			   &j.Department.Id,
																			   &j.Department.Name)
	if checkQuery(err) {
		session := sessions.GetSession(req)
		session.Set("teacher", j)
		RenderJson(resp, j)
	} else {
		RenderError(authFailed, resp)
	}
}

