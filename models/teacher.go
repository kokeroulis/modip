package models

import "github.com/kokeroulis/modip/db"

type Teacher struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
}

func (t *Teacher) Login(username string, password string) bool {
	var auth bool
	query := `SELECT id, name, email, departmentId, departmentName, auth
			  FROM teacher_auth($1, $2)`
	err := Db.Database.QueryRow(query, username, password).
		Scan(&t.Id, &t.Name, &t.Email, &t.Department.Id, &t.Department.Name, &auth)

	Db.CheckQuery(err, query)

	return auth
}

func (t *Teacher) Create(password string) bool {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM teacher_create($1, $2, $3, $4)`
	err := Db.Database.QueryRow(query, t.Name, password, t.Email, t.Department.Id).
		Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	return alreadyExists
}
