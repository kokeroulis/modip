package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry5 struct {
	// Διδακτικά Βοηθήματα
	Field1 string `schema:"lesson_didaktika_boi8imata_boi8imata_pou_dianemontai_stous_foithtes"`
	Field2 string `schema:"lesson_didaktika_boi8imata_epikairopoihsh_ton_boh8imaton"`
	Field3 string `schema:"lesson_didaktika_boi8imata_pososto_didaskomenis_ylis_pou_kaliptete"`
	Field4 string `schema:"lesson_didaktika_boi8imata_prostheti_bibliografia"`
	Field5 string `schema:"lesson_didaktika_boi8imata_gnostopoihsh_ylis_stous_foithtes"`
}

func (f *LessonCreateReportFormEntry5) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry5 SET
			lesson_didaktika_boi8imata_boi8imata_pou_dianemontai_stous_foithtes = $1,
			lesson_didaktika_boi8imata_epikairopoihsh_ton_boh8imaton = $2,
			lesson_didaktika_boi8imata_pososto_didaskomenis_ylis_pou_kaliptete = $3,
			lesson_didaktika_boi8imata_prostheti_bibliografia = $4,
			lesson_didaktika_boi8imata_gnostopoihsh_ylis_stous_foithtes = $5
			 WHERE lesson = $6 AND akademic_year = $7`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		f.Field3,
		f.Field4,
		f.Field5,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry5) Load(lessonId int, akademicYearId int) {
	query := `SELECT
			lesson_didaktika_boi8imata_boi8imata_pou_dianemontai_stous_foithtes,
			lesson_didaktika_boi8imata_epikairopoihsh_ton_boh8imaton,
			lesson_didaktika_boi8imata_pososto_didaskomenis_ylis_pou_kaliptete,
			lesson_didaktika_boi8imata_prostheti_bibliografia,
			lesson_didaktika_boi8imata_gnostopoihsh_ylis_stous_foithtes
			FROM LessonCreateReportFormEntry5
			WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5)

	Db.CheckQueryWithNoRows(err, query)
}
