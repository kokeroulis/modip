package models

import (
	"github.com/kokeroulis/modip/db"
)

type Lesson struct {
	Id            int    `schema:"id"`
	Name          string `schema:"name"`
	Department    int    `schema:"department"`
	IsPostDegree  bool   `schema:"is_post_degree"`
	CourseCode    string `schema:"course_code"`
	CardisoftCode string `schema:"cardisoft_code"`
}

func (l *Lesson) Create() {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM lesson_create($1, $2, $3, $4, $5)`

	err := Db.Database.QueryRow(query, l.Name, l.Department,
								l.IsPostDegree, l.CourseCode, l.CardisoftCode).
		Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	if alreadyExists {
		panic("The lesson already exists!!")
	}
}

func (l *Lesson) Load() {
	if l.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `SELECT courseCode, CardisoftCode, isPostDegree
			  FROM lesson
			  WHERE id = $1`
	err := Db.Database.QueryRow(query, l.Id).
		Scan(&l.CourseCode, &l.CardisoftCode, &l.IsPostDegree)

	Db.CheckQuery(err, query)
}

func (l *Lesson) Update() {
	if l.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `UPDATE lesson SET
			  courseCode = $1, cardisoftCode = $2,
			  isPostDegree = $3
			  WHERE id = $4`

	_, err := Db.Database.Exec(query,
								l.CourseCode,
								l.CardisoftCode,
								l.IsPostDegree)

	Db.CheckQueryWithNoRows(err, query)
}

func ListLessonsPreDegree() []*Department {
	return listLessons(false)
}

func ListLessonsPostDegree() []*Department {
	return listLessons(true)
}

func listLessons(postDegree bool) []*Department {
	var departmentList []*Department
	query := `SELECT id, name FROM department`

	rows, err := Db.Database.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := &Department{}
		if err := rows.Scan(&it.Id, &it.Name); err != nil {
			panic(err)
		} else {
			it.LoadLessons(postDegree)
			departmentList = append(departmentList, it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	return departmentList
}

