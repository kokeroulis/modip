package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry15 struct {
	// Η άποψη των σπουδαστών για το μάθημα
	Field1 string `schema:"apopsi_spoudaston_gia_to_ma8ima_yparxei_diadikasia_aksiologisis_tou_ma8imatos"`
	Field2 string `schema:"apopsi_spoudaston_gia_to_ma8ima_pws_aksiopoiountai_auta_ta_apotelesmata"`
}

func (f *LessonCreateReportFormEntry15) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry15 SET
		apopsi_spoudaston_gia_to_ma8ima_yparxei_diadikasia_aksiologisis_tou_ma8imatos = $1,
		apopsi_spoudaston_gia_to_ma8ima_pws_aksiopoiountai_auta_ta_apotelesmata = $2
		 WHERE lesson = $3 AND akademic_year = $4`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}
func (f *LessonCreateReportFormEntry15) Load(lessonId int, akademicYearId int) {
	query := `SELECT
		apopsi_spoudaston_gia_to_ma8ima_yparxei_diadikasia_aksiologisis_tou_ma8imatos,
		apopsi_spoudaston_gia_to_ma8ima_pws_aksiopoiountai_auta_ta_apotelesmata
		FROM LessonCreateReportFormEntry15
		WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2)

	Db.CheckQueryWithNoRows(err, query)
}
