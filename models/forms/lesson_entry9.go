package forms

import (
         "github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry9 struct {
	// Αξιολόγηση της επίδοσης των μαθητών στο μάθημα
	Field1 bool `schema:"lesson_aksiologisi_ths_epidosis_troipoi_aksiologisis"`
	Field2 string `schema:"lesson_aksiologisi_ths_epidosis_epeksigiseis_alla"`
	Field3 bool `schema:"lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg"`
	Field4 bool `schema:"lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou"`
	Field5 string `schema:"lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis"`
}

func (f *LessonCreateReportFormEntry9) Update(lessonId int, akademicYearId int) {
query := `UPDATE LessonCreateReportFormEntry9 SET
        lesson_aksiologisi_ths_epidosis_troipoi_aksiologisis = $1
        lesson_aksiologisi_ths_epidosis_epeksigiseis_alla = $2,
        lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg = $3,
        lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou = $4,
        lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis = $5
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
func (f *LessonCreateReportFormEntry9) Load(lessonId int, akademicYearId int) {
query := `SELECT
        lesson_aksiologisi_ths_epidosis_troipoi_aksiologisis,
        lesson_aksiologisi_ths_epidosis_alla,
        lesson_aksiologisi_ths_epidosis_epeksigiseis_alla,
        lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg,
        lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou,
        lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis
        FROM LessonCreateReportFormEntry9
        WHERE lesson = $1 AND akademic_year = $2`

err := Db.Database.QueryRow(query, lessonId, akademicYearId).
        Scan(&f.Field1,
                &f.Field2,
                &f.Field3,
                &f.Field4,
                &f.Field5)


Db.CheckQueryWithNoRows(err, query)
}
