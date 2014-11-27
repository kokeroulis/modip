package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry16 struct {
	// Σχόλια
	Field1 string   `schema:"sxolia_sxolia"`
}


=============== Start create update func
func (f *LessonCreateReportFormEntry16) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry16 SET
	sxolia_sxolia = $1
WHERE lesson = $2`

_, err := Db.Database.Exec(query,
	Field1
    lessonId)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *LessonCreateReportFormEntry16) Load(lessonId int) {
query := `SELECT
	sxolia_sxolia
FROM LessonCreateReportFormEntry16
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.Field1
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
