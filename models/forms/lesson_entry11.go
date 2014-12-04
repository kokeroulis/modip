package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry11 struct {
	// Αξιοποίηση Τεχνολογιών Πληροφορικής και Επικοινωνιών (ΤΠΕ)
	Field1 string `schema:"aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPE_sth_didaskalia_tou_ma8imatos"`
	Field2 string `schema:"aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_ma8isiaka_boh8imata_basismena_se_TPE"`
	Field3 string `schema:"aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPE_sthn_ergasthriaki_ekpedeusi"`
	Field4 string `schema:"aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPE_sthn_aksiologisi_ton_spoudaston_pws"`
	Field5 string `schema:"aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPT_sthn_epikoinonia_sas_me_tous_spoudastes_pws"`
}

func (f *LessonCreateReportFormEntry11) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry11 SET
			TPE_sth_didaskalia_tou_ma8imatos = $1,
			ma8isiaka_boh8imata_basismena_se_TPE = $2,
			TPE_sthn_ergasthriaki_ekpedeusi = $3,
			TPE_sthn_aksiologisi_ton_spoudaston_pws = $4,
			TPT_sthn_epikoinonia_sas_me_tous_spoudastes_pws = $5
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

func (f *LessonCreateReportFormEntry11) Load(lessonId int, akademicYearId int) {
	query := `SELECT
		TPE_sth_didaskalia_tou_ma8imatos,
		ma8isiaka_boh8imata_basismena_se_TPE,
		TPE_sthn_ergasthriaki_ekpedeusi,
		TPE_sthn_aksiologisi_ton_spoudaston_pws,
		TPT_sthn_epikoinonia_sas_me_tous_spoudastes_pws
		FROM LessonCreateReportFormEntry11
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
