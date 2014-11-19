package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry7 struct {
	// Επικοινωνία & Καθοδήγηση Σπουδαστών / Συνεργασίες
	Field1 string `schema:"lesson_epikoinonia_wres_grafeiou"`
	Field2 string `schema:"lesson_epikoinonia_ekepedeusi_methodos"`
	Field3 string `schema:"lesson_epikoinonia_organosi_ekpedeutikon_episkepseon"`
}

=============== Start create table
CREATE TABLE LessonCreateReportFormEntry7 (
	lesson int references lesson(id) on delete cascade,
	lesson_epikoinonia_wres_grafeiou text,
	lesson_epikoinonia_ekepedeusi_methodos text,
	lesson_epikoinonia_organosi_ekpedeutikon_episkepseon text
)
=============== End Create table
=============== Start create update func
func (f *LessonCreateReportFormEntry7) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry7 SET 
	lesson_epikoinonia_wres_grafeiou = $1,
	lesson_epikoinonia_ekepedeusi_methodos = $2,
	lesson_epikoinonia_organosi_ekpedeutikon_episkepseon = $3
WHERE lesson = $`

_, err := Db.Database.Exec(query, 
	Field1, 
	Field2, 
	Field3
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *LessonCreateReportFormEntry7) Load(lessonId int) {
query := `SELECT 
	lesson_epikoinonia_wres_grafeiou,
	lesson_epikoinonia_ekepedeusi_methodos,
	lesson_epikoinonia_organosi_ekpedeutikon_episkepseon
FROM LessonCreateReportFormEntry7
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO). 
	Scan(&f.Field1,
		&f.Field2,
		&f.Field3
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
