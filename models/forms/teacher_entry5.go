package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry5 struct {
	// Σύνδεση με την κοινωνία
	Field1 string        `schema:"sundesi_me_thn_koinonia_sundesi_me_tin_koinonia"`
}

=============== Start create table
CREATE TABLE  (
	lesson int references lesson(id) on delete cascade,
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia text
)
=============== End Create table
=============== Start create update func
func (f *) Update(lessonId int) {
query := `UPDATE  SET
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia = $1
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
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia
FROM
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.Field1
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
