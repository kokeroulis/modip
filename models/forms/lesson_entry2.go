package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry2 struct {
	// Διδασκαλίας
	Field1  string `schema:"lesson_didaskalia_typos_ma8iomatos1"`
	Field2  string `schema:"lesson_didaskalia_typos_ma8iomatos2"`
	Field3  string `schema:"lesson_didaskalia_typos_ma8iomatos3"`
	Field4  string `schema:"lesson_didaskalia_typos_ma8iomatos4"`
	Field5  string `schema:"lesson_didaskalia_typos_ma8iomatos5"`
}

func (f *LessonCreateReportFormEntry2) Update (lessonId int) {
	query := `UPDATE LessonCreateReportFormEntry2 SET
			 lesson_didaskalia_typos_ma8iomatos1 = $1,
			 lesson_didaskalia_typos_ma8iomatos2 = $2,
			 lesson_didaskalia_typos_ma8iomatos3 = $3,
			 lesson_didaskalia_typos_ma8iomatos4 = $4,
			 lesson_didaskalia_typos_ma8iomatos5 = $5
			 WHERE lesson = $6`

	_, err := Db.Database.Exec(query,
								f.Field1,
								f.Field2,
								f.Field3,
								f.Field4,
								f.Field5,
								lessonId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry2) Load (lessonId int) {
query := `SELECT
		 lesson_didaskalia_typos_ma8iomatos1,
		 lesson_didaskalia_typos_ma8iomatos2,
		 lesson_didaskalia_typos_ma8iomatos3,
		 lesson_didaskalia_typos_ma8iomatos4,
		 lesson_didaskalia_typos_ma8iomatos5
		 FROM LessonCreateReportFormEntry2
		 WHERE lesson = $1`

	err := Db.Database.QueryRow(query, lessonId).
		Scan(&f.Field1,
		 &f.Field2,
		 &f.Field3,
		 &f.Field4,
		 &f.Field5)

	Db.CheckQueryWithNoRows(err, query)
}
