package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry6 struct {
        // Aλλες εκπαιδευτικές δραστηριότητες;
	Field1 string `schema:"lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites"`
}

=============== Start create table
CREATE TABLE  (
	lesson int references lesson(id) on delete cascade,
	lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites text
)
=============== End Create table
=============== Start create update func
func (f *) Update(lessonId int) {
query := `UPDATE  SET
	lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites = $1
WHERE lesson = $`

_, err := Db.Database.Exec(query,
	Field1
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *) Load(lessonId int) {
query := `SELECT
	lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites
FROM
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.Field1
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
