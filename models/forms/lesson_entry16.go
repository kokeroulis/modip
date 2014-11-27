package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry16 struct {
	// Σχόλια
	Field1 string `schema:"sxolia_sxolia"`
}

func (f *LessonCreateReportFormEntry16) Update(lessonId int) {
	query := `UPDATE LessonCreateReportFormEntry16 SET
			sxolia_sxolia = $1
			WHERE lesson = $2`

	_, err := Db.Database.Exec(query,
		f.Field1,
		lessonId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry16) Load(lessonId int) {
	query := `SELECT
		sxolia_sxolia
		FROM LessonCreateReportFormEntry16
		WHERE lesson = $1`

	err := Db.Database.QueryRow(query, lessonId).
		Scan(&f.Field1)

	Db.CheckQueryWithNoRows(err, query)
}
