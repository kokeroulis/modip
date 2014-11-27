package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry2 struct {
	// Διδασκαλίας
	Field1  string `schema:"lesson_didaskalia_typos_ma8iomatos1"`
	Field2  string `schema:"lesson_didaskalia_typos_ma8iomatos2"`
	Field3  string `schema:"lesson_didaskalia_typos_ma8iomatos3"`
	Field4  string `schema:"lesson_didaskalia_typos_ma8iomatos4"`
	Field5  string `schema:"lesson_didaskalia_typos_ma8iomatos5"`
    Field6 string  `schema:"lesson_didaskalia_wres_didaskaliasi_8"`
    Field7 string  `schema:"lesson_didaskalia_wres_didaskalias_AP"`
    Field8 string  `schema:"lesson_didaskalia_wres_didaskalias_E"`
    Field9 string  `schema:"lesson_didaskalia_monadesECTS_8AP"`
    Field10 string `schema:"lesson_didaskalia_monadesECTS_E"`
    Field11 string `schema:"lesson_didaskalia_wres_didaskalias_ana_bdomada"`
    Field12 bool   `schema:"lesson_didaskalia_wres_ergasthriou_askhshs"`
    Field13 string `schema:"lesson_didaskalia_wres_ergasthriou_ana_bdomada"`
    Field14 string `schema:"lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino"`
    Field15 string `schema:"lesson_didaskalia_problepomenes_wres_didaskalias_mikron_omadon_ana_eksamino"`
    Field16 string `schema:"lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli"`
    Field17 bool   `schema:"lesson_didaskalia_pollapli_bibiografia"`
    Field18 bool   `schema:"lesson_didaskalia_diati8ete_ergasia_proodos"`
    Field19 bool   `schema:"lesson_didaskalia_ypoxreotiki_ergasia_proodos"`
}

func (f *LessonCreateReportFormEntry2) Update(lessonId int) {
	query := `UPDATE LessonCreateReportFormEntry2 SET
			 lesson_didaskalia_typos_ma8iomatos1 = $1,
			 lesson_didaskalia_typos_ma8iomatos2 = $2,
			 lesson_didaskalia_typos_ma8iomatos3 = $3,
			 lesson_didaskalia_typos_ma8iomatos4 = $4,
			 lesson_didaskalia_typos_ma8iomatos5 = $5
			 WHERE lesson = $6`

	_, err := Db.Database.Exec(query,
								f.Field1,
								f.Field2,
								f.Field3,
								f.Field4,
								f.Field5,
								lessonId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry2) Load(lessonId int) {
query := `SELECT
		 lesson_didaskalia_typos_ma8iomatos1,
		 lesson_didaskalia_typos_ma8iomatos2,
		 lesson_didaskalia_typos_ma8iomatos3,
		 lesson_didaskalia_typos_ma8iomatos4,
		 lesson_didaskalia_typos_ma8iomatos5
		 FROM LessonCreateReportFormEntry2
		 WHERE lesson = $1`

	err := Db.Database.QueryRow(query, lessonId).
		Scan(&f.Field1,
		 &f.Field2,
		 &f.Field3,
		 &f.Field4,
		 &f.Field5)

	Db.CheckQueryWithNoRows(err, query)
}
