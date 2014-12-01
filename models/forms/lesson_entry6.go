package forms

import (
         "github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry6 struct {
        // Aλλες εκπαιδευτικές δραστηριότητες;
        Field1 string `schema:"lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites"`
}

func (f *LessonCreateReportFormEntry6) Update(lessonId int) {
    query := `UPDATE LessonCreateReportFormEntry6 SET
        lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites = $1
        WHERE lesson = $2`

    _, err := Db.Database.Exec(query,
        f.Field1,
        lessonId)

    Db.CheckQueryWithNoRows(err, query)
}
func (f *LessonCreateReportFormEntry6) Load(lessonId int) {
    query := `SELECT
                lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites
                FROM LessonCreateReportFormEntry6
                WHERE lesson = $2`

    err := Db.Database.QueryRow(query, lessonId).
                Scan(&f.Field1)

    Db.CheckQueryWithNoRows(err, query)
}
