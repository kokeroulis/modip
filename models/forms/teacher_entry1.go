package forms

import (
	"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry1 struct {
	// Αριθμός Δημοσιεύσεων
	Id      int    `schema:-`
	Field1  string `schema:"author"`
	Field2  string `schema:"title"`
	Field3  bool   `schema:"is_magazine"`
	Field4  string `schema:"publisher"`
	Field5  string `schema:"publication_date"`
	Field6  string `schema:"type"`
}

func (f *TeacherCreateReportFormEntry1) Create(teacherId int) {
	query := `INSERT INTO TeacherCreateReportFormEntry1
			  (author, title, is_magazine, publisher,
			   publication_date, type, teacher)
			  VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err := Db.Database.Exec(query,
								f.Field1,
								f.Field2,
								f.Field3,
								f.Field4,
								f.Field5,
								f.Field6,
								teacherId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *TeacherCreateReportFormEntry1) Update() {
	query := `UPDATE TeacherCreateReportFormEntry1
			  SET author = $1, title = $2, is_magazine = $3,
			  publisher = $4,
			  publication_date = $5, type = $6
			  WHERE id = $2`

	err := Db.Database.QueryRow(query, f.Id).
		Scan(&f.Field1,
			 &f.Field2,
			 &f.Field3,
			 &f.Field4,
			 &f.Field5,
			 &f.Field6)

	Db.CheckQueryWithNoRows(err, query)
}

func ListAllTeacherCreateReportForm1(teacherId int) []TeacherCreateReportFormEntry1 {
	var formList []TeacherCreateReportFormEntry1

	query := `SELECT
			  id, author, title, is_magazine,
			  publisher, publication_date, type
			  FROM TeacherCreateReportFormEntry1
			  WHERE teacher = $1`

	rows, err := Db.Database.Query(query, teacherId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()


	for rows.Next() {
		it := TeacherCreateReportFormEntry1{}

		if err := rows.Scan(&it.Id,
							&it.Field1,
							&it.Field2,
							&it.Field3,
							&it.Field4,
							&it.Field5,
							&it.Field6); err != nil {
			panic(err)
		} else {
			formList = append(formList, it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	return formList
}
