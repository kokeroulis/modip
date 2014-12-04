package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry4 struct {
	// Διδακτέα Ύλη
	Field1 string `schema:"lesson_didaktea_yli_anaprosarmogi_mathimatos"`
	Field2 string `schema:"lesson_didaktea_yli_epikalipsi_ylis_me_alla_ma8imata"`
}

func (f *LessonCreateReportFormEntry4) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry4 SET
			  lesson_didaktea_yli_anaprosarmogi_mathimatos = $1,
			  lesson_didaktea_yli_epikalipsi_ylis_me_alla_ma8imata = $2
			   WHERE lesson = $3 AND akademic_year = $4`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry4) Load(lessonId int, akademicYearId int) {
	query := `SELECT
			 lesson_didaktea_yli_anaprosarmogi_mathimatos,
			 lesson_didaktea_yli_epikalipsi_ylis_me_alla_ma8imata
			 FROM LessonCreateReportFormEntry4
			 WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2)

	Db.CheckQueryWithNoRows(err, query)
}
