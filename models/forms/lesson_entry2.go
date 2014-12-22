package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry2 struct {
	// Διδασκαλίας
	Field1  string `schema:"lesson_didaskalia_typos_ma8iomatos"`
	Field2  string `schema:"lesson_didaskalia_wres_didaskaliasi_8"`
	Field3  string `schema:"lesson_didaskalia_wres_didaskalias_AP"`
	Field4  string `schema:"lesson_didaskalia_wres_didaskalias_E"`
	Field5  string `schema:"lesson_didaskalia_monadesECTS_8AP"`
	Field6 string `schema:"lesson_didaskalia_monadesECTS_E"`
	Field7 string `schema:"lesson_didaskalia_wres_didaskalias_ana_bdomada"`
	Field8 bool   `schema:"lesson_didaskalia_wres_ergasthriou_askhshs"`
	Field9 string `schema:"lesson_didaskalia_wres_ergasthriou_ana_bdomada"`
	Field10 string `schema:"lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino"`
	Field11 string `schema:"lesson_didaskalia_problepomenes_wres_8ewrias_ana_eksamino"`
	Field12 string `schema:"lesson_didaskalia_problepomenes_wres_didaskalias_mikron_omadon_ana_eksamino"`
	Field13 string `schema:"lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli"`
	Field14 bool   `schema:"lesson_didaskalia_pollapli_bibiografia"`
	Field15 bool   `schema:"lesson_didaskalia_diati8ete_ergasia_proodos"`
	Field16 bool   `schema:"lesson_didaskalia_ypoxreotiki_ergasia_proodos"`
}

func (f *LessonCreateReportFormEntry2) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry2 SET
             lesson_didaskalia_typos_ma8iomatos = $ 1,
             lesson_didaskalia_wres_didaskaliasi_8 = $2,
             lesson_didaskalia_wres_didaskalias_AP = $3,
             lesson_didaskalia_wres_didaskalias_E = $4,
             lesson_didaskalia_monadesECTS_8AP = $5,
             lesson_didaskalia_monadesECTS_E = $6,
             lesson_didaskalia_wres_didaskalias_ana_bdomada = $7,
             lesson_didaskalia_wres_ergasthriou_askhshs = $8,
             lesson_didaskalia_wres_ergasthriou_ana_bdomada = $9,
             lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino = $10,
             lesson_didaskalia_problepomenes_wres_8ewrias_ana_eksamino = $11,
             lesson_didaskalia_problepomenes_wres_didaskalias_mikron_omadon_ana_eksamino = $12,
             lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli = $13,
             lesson_didaskalia_pollapli_bibiografia = $14,
             lesson_didaskalia_diati8ete_ergasia_proodos = $15,
             lesson_didaskalia_ypoxreotiki_ergasia_proodos = $16
			  WHERE lesson = $17 AND akademic_year = $18`

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
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry2) Load(lessonId int, akademicYearId int) {
	query := `SELECT
             lesson_didaskalia_typos_ma8iomatos,
             lesson_didaskalia_wres_didaskaliasi_8,
             lesson_didaskalia_wres_didaskalias_AP,
             lesson_didaskalia_wres_didaskalias_E,
             lesson_didaskalia_monadesECTS_8AP,
             lesson_didaskalia_monadesECTS_E,
             lesson_didaskalia_wres_didaskalias_ana_bdomada,
             lesson_didaskalia_wres_ergasthriou_askhshs,
             lesson_didaskalia_wres_ergasthriou_ana_bdomada,
             lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino,
             lesson_didaskalia_problepomenes_wres_8ewrias_ana_eksamino,
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
		&f.Field16)

	Db.CheckQueryWithNoRows(err, query)
}
