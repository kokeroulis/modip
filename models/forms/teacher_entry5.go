package forms

import (
	"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry5 struct {
	// Σύνδεση με την κοινωνία
	Field1 string `schema:"sundesi_me_thn_koinonia_sundesi_me_tin_koinonia"`
}

func (f *TeacherCreateReportFormEntry5) Update(teacherId int, akademicYearId int) {
	query := `UPDATE TeacherCreateReportFormEntry5
			SET sundesi_me_thn_koinonia_sundesi_me_tin_koinonia = $1
			 WHERE teacher = $2 AND akademic_year = $3`

	_, err := Db.Database.Exec(query,
		f.Field1,
		teacherId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *TeacherCreateReportFormEntry5) Load(teacherId int, akademicYearId int) {
	query := `SELECT
			sundesi_me_thn_koinonia_sundesi_me_tin_koinonia
			FROM TeacherCreateReportFormEntry5
			WHERE teacher = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, teacherId,
		akademicYearId).
		Scan(&f.Field1)

	Db.CheckQueryWithNoRows(err, query)
}
