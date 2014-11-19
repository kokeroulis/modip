package forms

import (
	_"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry8 struct {
	// Συμμετοχή Σπουδαστών
	Field1 string `schema:"lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi"`
	Field2 int `schema:"lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi"`
}

=============== Start create table
CREATE TABLE LessonCreateReportFormEntry8 (
	lesson int references lesson(id) on delete cascade,
	lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi text,
	lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi int
)
=============== End Create table
=============== Start create update func
func (f *LessonCreateReportFormEntry8) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry8 SET 
	lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi = $1,
	lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi = $2
WHERE lesson = $`

_, err := Db.Database.Exec(query, 
	Field1, 
	Field2
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *LessonCreateReportFormEntry8) Load(lessonId int) {
query := `SELECT 
	lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi,
	lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi
FROM LessonCreateReportFormEntry8
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO). 
	Scan(&f.Field1,
		&f.Field2
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
