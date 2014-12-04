package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry13 struct {
	// Στατιστικά Στοιχεία Σπουδαστών
	Field1 string `schema:"statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima"`
}

func (f *LessonCreateReportFormEntry13) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry13 SET
			statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima = $1
		 WHERE lesson = $2 AND akademic_year = $3`

	_, err := Db.Database.Exec(query,
		f.Field1,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry13) Load(lessonId int, akademicYearId int) {
	query := `SELECT
			statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima
		FROM LessonCreateReportFormEntry13
		WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1)

	Db.CheckQueryWithNoRows(err, query)
}
