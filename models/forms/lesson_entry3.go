package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry3 struct {
	// Ενημέρωση _ Αξιολόγηση
	Field1 int    `schema:"lesson_enhmerwsi_selida_odigou_spoudon"`
	Field2 string `schema:"lesson_enhmerwsi_website"`
	Field3 bool   `schema:"lesson_enhmerwsi_aksiologisi_E"`
	Field4 bool   `schema:"lesson_enhmerwsi_aksiologisi_8"`
	Field5 int    `schema:"lesson_enhmerwsi_ari8mos_spoudaston_pou_aksiologisan_to_ma8ima"`
}

func (f *LessonCreateReportFormEntry3) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry3 SET
			 lesson_enhmerwsi_selida_odigou_spoudon = $1,
			 lesson_enhmerwsi_website = $2,
			 lesson_enhmerwsi_aksiologisi_E = $3,
			 lesson_enhmerwsi_aksiologisi_8 = $4,
			 lesson_enhmerwsi_ari8mos_spoudaston_pou_aksiologisan_to_ma8ima = $5
			  WHERE lesson = $6 AND akademic_year = $7`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		f.Field3,
		f.Field4,
		f.Field5,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry3) Load(lessonId int, akademicYearId int) {
	query := `SELECT
			 lesson_enhmerwsi_selida_odigou_spoudon,
			 lesson_enhmerwsi_website,
			 lesson_enhmerwsi_aksiologisi_E,
			 lesson_enhmerwsi_aksiologisi_8,
			 lesson_enhmerwsi_ari8mos_spoudaston_pou_aksiologisan_to_ma8ima
			 FROM LessonCreateReportFormEntry3
			 WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5)

	Db.CheckQueryWithNoRows(err, query)
}
