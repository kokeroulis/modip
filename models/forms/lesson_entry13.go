package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry13 struct {
	// Στατιστικά Στοιχεία Σπουδαστών
	Field1 string  `schema:"statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima"`
}

=============== Start create update func
func (f *LessonCreateReportFormEntry13) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry13 SET
	statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima = $1
WHERE lesson = $`

_, err := Db.Database.Exec(query,
	Field1
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *LessonCreateReportFormEntry13) Load(lessonId int) {
query := `SELECT
	statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima
FROM LessonCreateReportFormEntry13
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.Field1
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
