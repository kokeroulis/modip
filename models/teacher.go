package models

import "github.com/kokeroulis/modip/db"

type Teacher struct {
	Id         int        `schema:"id"`
	Name       string     `schema:"name"`
	Email      string     `schema:"email"`
	Department Department `schema:"department"`
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

func (t *Teacher) Load() {
	if t.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `SELECT name, email
			  FROM teacher
			  WHERE id = $1`
	err := Db.Database.QueryRow(query, t.Id).
		Scan(&t.Name, &t.Email)

	Db.CheckQuery(err, query)
}

func (t *Teacher) Update() {
	if t.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `UPDATE teacher SET
			  name = $1, email = $2
			  WHERE id = $3`

	_, err := Db.Database.Exec(query,
								t.Name,
								t.Email,
								t.Id)

	Db.CheckQueryWithNoRows(err, query)
}

