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

type TeacherInfo struct {
	Teacher Teacher           `json:"teacher"`
	Books   []BookOrPaperInfo `json:"books"`
	Papers  []BookOrPaperInfo `json:"papers"`
}

func (t *Teacher) retrieveInfo(query string, channel chan []BookOrPaperInfo) {
	var args []interface{}

	args = append(args, t.Id)


	result := Db.SqlRowsToArray(query, args, &BookOrPaperInfo{})
	l := []BookOrPaperInfo{}
	for _, it := range result {
		v := it.(*BookOrPaperInfo)
		l = append(l, *v)
	}

	channel <- l
}

func (t *Teacher) Info() TeacherInfo {
	books := make(chan []BookOrPaperInfo)
	booksQuery := `SELECT id, title FROM book WHERE teacher = $1`

	papers := make(chan []BookOrPaperInfo)
	papersQuery := `SELECT id, title FROM paper WHERE teacher = $1`

	go t.retrieveInfo(booksQuery, books)
	go t.retrieveInfo(papersQuery, papers)

	var args []interface{}

	args = append(args, t.Id)

	info := TeacherInfo{
		Teacher: *t,
		Books: <-books,
		Papers: <-papers,
	}

	return info
}
