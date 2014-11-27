package forms

import (
	"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry5 struct {
	// Σύνδεση με την κοινωνία
	Field1 string `schema:"sundesi_me_thn_koinonia_sundesi_me_tin_koinonia"`
}

func (f *TeacherCreateReportFormEntry5) Update(teacherId int) {
	query := `UPDATE  SET
			sundesi_me_thn_koinonia_sundesi_me_tin_koinonia = $1
			WHERE teacher = $2`

	_, err := Db.Database.Exec(query,
		f.Field1,
		teacherId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *TeacherCreateReportFormEntry5) Load(teacherId int) {
	query := `SELECT
			sundesi_me_thn_koinonia_sundesi_me_tin_koinonia
			FROM
			WHERE teacher = $1`

	err := Db.Database.QueryRow(query, teacherId).
		Scan(&f.Field1)

	Db.CheckQueryWithNoRows(err, query)
}
