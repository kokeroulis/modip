package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry8 struct {
	// Συμμετοχή Σπουδαστών

	Field1 string `schema:"lesson_summetoxi_spoudaston_sto_mathima"`
	Field2 string `schema:"lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi"`
	Field3 int `schema:"lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi"`
}

func (f *LessonCreateReportFormEntry8) Update(lessonId int, akademicYearId int) {
query := `UPDATE LessonCreateReportFormEntry8 SET
	lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi = $1,
	lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi = $2
WHERE lesson = $3 AND akademic_year = $4`

_, err := Db.Database.Exec(query,
	f.Field1,
	f.Field2,
    lessonId,
    akademicYearId)

Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry8) Load(lessonId int, akademicYearId int) {
query := `SELECT
	lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi,
	lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi
    FROM LessonCreateReportFormEntry8
    WHERE lesson = $1 AND akademic_year = $2`

err := Db.Database.QueryRow(query, lessonId, akademicYearId).
	Scan(&f.Field1,
		&f.Field2)

Db.CheckQueryWithNoRows(err, query)
}
