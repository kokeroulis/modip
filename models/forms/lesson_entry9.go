package forms

import (
         "github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry9 struct {
	// Αξιολόγηση της επίδοσης των μαθητών στο μάθημα 
	Field1 bool `schema:"lesson_aksiologisi_ths_epidosis_grapti_eksetasi"`
	Field2 bool `schema:"lesson_aksiologisi_ths_epidosis_proforiki_eksetasi"`
	Field3 bool `schema:"lesson_aksiologisi_ths_epidosis_proodos"`
	Field4 bool `schema:"lesson_aksiologisi_ths_epidosis_katoi_kon_ergasia"`
	Field5 bool `schema:"lesson_aksiologisi_ths_epidosis_proforiki_parousiasi_ergasias"`
	Field6 bool `schema:"lesson_aksiologisi_ths_epidosis_ergashtirio_praktikes_askiseis"`
	Field7 bool `schema:"lesson_aksiologisi_ths_epidosis_alla"`
	Field8 string `schema:"lesson_aksiologisi_ths_epidosis_epeksigiseis_alla"`
	Field9 bool `schema:"lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg"`
	Field10 bool `schema:"lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou"`
	Field11 string `schema:"lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis"`
}

func (f *LessonCreateReportFormEntry9) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry9 SET
        lesson_aksiologisi_ths_epidosis_grapti_eksetasi = $1,
        lesson_aksiologisi_ths_epidosis_proforiki_eksetasi = $2,
        lesson_aksiologisi_ths_epidosis_proodos = $3,
        lesson_aksiologisi_ths_epidosis_katoi_kon_ergasia = $4,
        lesson_aksiologisi_ths_epidosis_proforiki_parousiasi_ergasias = $5,
        lesson_aksiologisi_ths_epidosis_ergashtirio_praktikes_askiseis = $6,
        lesson_aksiologisi_ths_epidosis_alla = $7,
        lesson_aksiologisi_ths_epidosis_epeksigiseis_alla = $8,
        lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg = $9,
        lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou = $10,
        lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis = $11
WHERE lesson = $12`

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
        lessonId)

Db.CheckQueryWithNoRows(err, query)
}
func (f *LessonCreateReportFormEntry9) Load(lessonId int) {
query := `SELECT
        lesson_aksiologisi_ths_epidosis_grapti_eksetasi,
        lesson_aksiologisi_ths_epidosis_proforiki_eksetasi,
        lesson_aksiologisi_ths_epidosis_proodos,
        lesson_aksiologisi_ths_epidosis_katoi_kon_ergasia,
        lesson_aksiologisi_ths_epidosis_proforiki_parousiasi_ergasias,
        lesson_aksiologisi_ths_epidosis_ergashtirio_praktikes_askiseis,
        lesson_aksiologisi_ths_epidosis_alla,
        lesson_aksiologisi_ths_epidosis_epeksigiseis_alla,
        lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg,
        lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou,
        lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis
FROM LessonCreateReportFormEntry9
WHERE lesson = $1`

err := Db.Database.QueryRow(query, lessonId).
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
                &f.Field11)

Db.CheckQueryWithNoRows(err, query)
}
