package forms

import (
	_ "github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry5 struct {
	// Σύνδεση με την κοινωνία
	Field1 string `schema:"sundesi_me_thn_koinonia_sundesi_me_tin_koinonia"`
}

func (f *TeacherCreateReportFormEntry5) Update(lessonId int) {
	query := `UPDATE  SET
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia = $1
WHERE lesson = $2`

	_, err := Db.Database.Exec(query,
		Field1,
		lessonId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *TeacherCreateReportFormEntry5) Load(lessonId int) {
	query := `SELECT
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia
FROM
WHERE lesson = $1`

	err := Db.Database.QueryRow(query, lessonId).
		Scan(&f.Field1)

	Db.CheckQueryWithNoRows(err, query)
}
