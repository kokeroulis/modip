package models

import (
	"github.com/kokeroulis/modip/db"
)

type Lesson struct {
	Id           int
	Name         string
	Department   int
	IsPostDegree bool
	CourseCode   string
}

func (l *Lesson) Create() {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM lesson_create($1, $2, $3, $4, $5)`

	err := Db.Database.QueryRow(query, l.Id, l.Name, l.Department, l.IsPostDegree, c.CourseCode).
		Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	if alreadyExists {
		panic("The lesson already exists!!")
	}
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

