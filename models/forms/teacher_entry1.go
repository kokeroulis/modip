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
			  (author, title, is_magazine, publisher, publication_datei, type)
			  VALUES($1, $2, $3, $4, $5, $6)
			  WHERE teacher = $3`

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

func (f *TeacherCreateReportFormEntry1) Update(teacherId int) {
	query := `UPDATE TeacherCreateReportFormEntry1
			  SET author = $1, title = $2, is_magazine $3,
			  publisher = $4,
			  publication_date = $5, type = $5
			  WHERE teacher = $1 AND id = $2`

	err := Db.Database.QueryRow(query, teacherId, f.Id).
		Scan(&f.Field1,
			 &f.Field2,
			 &f.Field3,
			 &f.Field4,
			 &f.Field5,
			 &f.Field6)

	Db.CheckQueryWithNoRows(err, query)
}

