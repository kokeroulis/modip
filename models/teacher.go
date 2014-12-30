package models

import "github.com/kokeroulis/modip/db"

type Teacher struct {
	Id         int        `schema:"id"`
	Name       string     `schema:"name"`
	Email      string     `schema:"email"`
	Department Department `schema:"department"`
    Username   string     `schema:"username"`
    Type       int        `schema:"type"`
    TypeName   string     //we only need it in the html
}

var TeacherType map[int]string

func init() {
    TeacherType = map[int]string{
        1: "aaa",
        2: "bbb",
        3: "ccc",
        4: "ddd",
        5: "eee",
    }
}

func (t *Teacher) Login(username string, password string) bool {
	var auth bool
	query := `SELECT id, name, email, departmentId, departmentName, auth, username, type
			  FROM teacher_auth($1, $2)`
	err := Db.Database.QueryRow(query, username, password).
		Scan(&t.Id, &t.Name, &t.Email, &t.Department.Id, &t.Department.Name, &auth, &t.Username, &t.Type)


    if auth {
        //populate our string type
        //we need it for the html
        t.TypeName = TeacherType[t.Type]
    }

	Db.CheckQuery(err, query)

	return auth
}

func (t *Teacher) Create(password string) bool {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM teacher_create($1, $2, $3, $4, $5, $6)`
	err := Db.Database.QueryRow(query, t.Name, password, t.Email, t.Department.Id, t.Username, t.Type).
		Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	return alreadyExists
}

func (t *Teacher) Load() {
	if t.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `SELECT name, email, username, type
			  FROM teacher
			  WHERE id = $1`
	err := Db.Database.QueryRow(query, t.Id).
		Scan(&t.Name, &t.Email, &t.Username, &t.Type)

    //populate our string type
    t.TypeName = TeacherType[t.Type]

	Db.CheckQuery(err, query)
}

func (t *Teacher) Update() {
	if t.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `UPDATE teacher SET
			  email = $1, name = $3,
              type = $4
			  WHERE id = $2`

	_, err := Db.Database.Exec(query,
								t.Email,
								t.Id,
                                t.Name,
                                t.Type)

	Db.CheckQueryWithNoRows(err, query)
}

