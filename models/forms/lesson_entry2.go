package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry2 struct {
	// Διδασκαλίας
	Field1  bool   `schema:"lesson_didaskalia_typos_ma8iomatos1"`
	Field2  bool   `schema:"lesson_didaskalia_typos_ma8iomatos2"`
	Field3  bool   `schema:"lesson_didaskalia_typos_ma8iomatos3"`
	Field4  bool   `schema:"lesson_didaskalia_typos_ma8iomatos4"`
	Field5  bool   `schema:"lesson_didaskalia_typos_ma8iomatos5"`
	Field6  string `schema:"lesson_didaskalia_wres_didaskaliasi_8"`
	Field7  string `schema:"lesson_didaskalia_wres_didaskalias_AP"`
	Field8  string `schema:"lesson_didaskalia_wres_didaskalias_E"`
	Field9  string `schema:"lesson_didaskalia_monadesECTS_8AP"`
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

func (f *LessonCreateReportFormEntry2) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry2 SET
             lesson_didaskalia_typos_ma8iomatos1 = $1,
             lesson_didaskalia_typos_ma8iomatos2 = $2,
             lesson_didaskalia_typos_ma8iomatos3 = $3,
             lesson_didaskalia_typos_ma8iomatos4 = $4,
             lesson_didaskalia_typos_ma8iomatos5 = $5,
             lesson_didaskalia_wres_didaskaliasi_8 = $6,
             lesson_didaskalia_wres_didaskalias_AP = $7,
             lesson_didaskalia_wres_didaskalias_E = $8,
             lesson_didaskalia_monadesECTS_8AP = $9,
             lesson_didaskalia_monadesECTS_E = $10,
             lesson_didaskalia_wres_didaskalias_ana_bdomada = $11,
             lesson_didaskalia_wres_ergasthriou_askhshs = $12,
             lesson_didaskalia_wres_ergasthriou_ana_bdomada = $13,
             lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino = $14,
             lesson_didaskalia_problepomenes_wres_didaskalias_mikron_omadon_ana_eksamino = $15,
             lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli = $16,
             lesson_didaskalia_pollapli_bibiografia = $17,
             lesson_didaskalia_diati8ete_ergasia_proodos = $18,
             lesson_didaskalia_ypoxreotiki_ergasia_proodos = $19
			  WHERE lesson = $20 AND akademic_year = $21`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		f.Field3,
		f.Field4,
		f.Field5,
		f.Field6,
		f.Field7,
		f.Field8,
		f.Field9,
		f.Field10,
		f.Field11,
		f.Field12,
		f.Field13,
		f.Field14,
		f.Field15,
		f.Field16,
		f.Field17,
		f.Field18,
		f.Field19,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry2) Load(lessonId int, akademicYearId int) {
	query := `SELECT
             lesson_didaskalia_typos_ma8iomatos1,
             lesson_didaskalia_typos_ma8iomatos2,
             lesson_didaskalia_typos_ma8iomatos3,
             lesson_didaskalia_typos_ma8iomatos4,
             lesson_didaskalia_typos_ma8iomatos5,
             lesson_didaskalia_wres_didaskaliasi_8,
             lesson_didaskalia_wres_didaskalias_AP,
             lesson_didaskalia_wres_didaskalias_E,
             lesson_didaskalia_monadesECTS_8AP,
             lesson_didaskalia_monadesECTS_E,
             lesson_didaskalia_wres_didaskalias_ana_bdomada,
             lesson_didaskalia_wres_ergasthriou_askhshs,
             lesson_didaskalia_wres_ergasthriou_ana_bdomada,
             lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino,
             lesson_didaskalia_problepomenes_wres_didaskalias_mikron_omadon_ana_eksamino,
             lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli,
             lesson_didaskalia_pollapli_bibiografia,
             lesson_didaskalia_diati8ete_ergasia_proodos,
             lesson_didaskalia_ypoxreotiki_ergasia_proodos
		 FROM LessonCreateReportFormEntry2
		 WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5,
		&f.Field6,
		&f.Field7,
		&f.Field8,
		&f.Field9,
		&f.Field10,
		&f.Field11,
		&f.Field12,
		&f.Field13,
		&f.Field14,
		&f.Field15,
		&f.Field16,
		&f.Field17,
		&f.Field18,
		&f.Field19)

	Db.CheckQueryWithNoRows(err, query)
}
