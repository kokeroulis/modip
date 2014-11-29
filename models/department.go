package models

import (
	"github.com/kokeroulis/modip/db"
)

type Department struct {
	Id       int
	Name     string
	Lessons  []*Lesson
	Teachers []*Teacher
}

func (d *Department) AddLesson(l *Lesson) {
	d.Lessons = append(d.Lessons, l)
}

func (d *Department) AddTeacher(t *Teacher) {
	d.Teachers = append(d.Teachers, t)
}

func (d *Department) LoadLessons(postDegree bool) {
	query := `SELECT id, name, courseCode, cardisoftCode, isPostDegree
			  FROM lesson
			  WHERE department = $1 AND isPostDegree = $2`

	rows, err := Db.Database.Query(query, d.Id, postDegree)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := &Lesson{
			IsPostDegree: postDegree,
		}

		if err := rows.Scan(&it.Id,
							&it.Name,
							&it.CourseCode,
							&it.CardisoftCode,
							&it.IsPostDegree); err != nil {
			panic(err)
		} else {
			d.AddLesson(it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}
}

func (d *Department) LoadStuff() {
	query := `SELECT id, name, email
			  FROM teacher
			  WHERE department = $1`

	rows, err := Db.Database.Query(query, d.Id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := &Teacher{}

		if err := rows.Scan(&it.Id,
							&it.Name,
							&it.Email); err != nil {
			panic(err)
		} else {
			d.AddTeacher(it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}
}

func ListAllDepartments(loadLessons bool) []Department {
	var departmentList []Department

	query := `SELECT id, name FROM department`

	rows, err := Db.Database.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := Department{}

		if err := rows.Scan(&it.Id, &it.Name); err != nil {
			panic(err)
		} else {
			departmentList = append(departmentList, it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	return departmentList
}

