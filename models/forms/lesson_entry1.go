package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry1 struct {
	// Περιγραφή
	Field1 string `schema:"lesson_perigrafi_periexomeno_ma8imatos"`
	Field2 string `schema:"lesson_perigrafi_ma8isiakoi_stoxoi"`
}

func (f *LessonCreateReportFormEntry1) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry1 SET
			  lesson_perigrafi_periexomeno_ma8imatos = $1,
			  lesson_perigrafi_ma8isiakoi_stoxoi = $2
			   WHERE lesson = $3 AND akademic_year = $4`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry1) Load(lessonId int, akademicYearId int) {
	query := `SELECT
			  lesson_perigrafi_periexomeno_ma8imatos,
			  lesson_perigrafi_ma8isiakoi_stoxoi
			  FROM LessonCreateReportFormEntry1
			  WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2)

	Db.CheckQueryWithNoRows(err, query)
}
