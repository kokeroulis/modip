package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportForm struct {
	Id           int    `json:"id"`

	// Περιγραφή
	Field1_1  string `schema:"lesson_perigrafi_periexomeno_ma8imatos"`
	Field1_2  string `schema:"lesson_perigrafi_ma8isiakoi_stoxoi"`

	// Διδασκαλίας
	Field2_1  string `schema:"lesson_didaskalia_typos_ma8iomatos1"`
	Field2_2  string `schema:"lesson_didaskalia_typos_ma8iomatos2"`
	Field2_3  string `schema:"lesson_didaskalia_typos_ma8iomatos3"`
	Field2_4  string `schema:"lesson_didaskalia_typos_ma8iomatos4"`
	Field2_5  string `schema:"lesson_didaskalia_typos_ma8iomatos5"`

	Field2_6  string `schema:"lesson_didaskalia_wres_didaskaliasi_8"`
	Field2_7  string `schema:"lesson_didaskalia_wres_didaskalias_AP"`
	Field2_8  string `schema:"lesson_didaskalia_wres_didaskalias_E"`
	Field2_9  string `schema:"lesson_didaskalia_monadesECTS_8AP"`
	Field2_10 string `schema:"lesson_didaskalia_monadesECTS_E"`
	Field2_11 string `schema:"lesson_didaskalia_wres_didaskalias_ana_bdomada"`
	Field2_12 string `schema:"lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino"`
	Field2_13 string `schema:"lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli"`
	Field2_14 bool   `schema:"lesson_didaskalia_pollapli_bibiografia"`
	Field2_15 bool   `schema:"lesson_didaskalia_diati8ete_ergasia_proodos"`
	Field2_16 bool   `schema:"lesson_didaskalia_ypoxreotiki_ergasia_proodos"`

	// Ενημέρωση _ Αξιολόγηση
	Field3_1  int    `schema:"lesson_enhmerwsi_selida_odigou_spoudon"`
	Field3_2  string `schema:"lesson_enhmerwsi_website"`
	Field3_3  bool   `schema:"lesson_enhmerwsi_aksiologisi_E"`
	Field3_4  bool   `schema:"lesson_enhmerwsi_aksiologisi_8"`
	Field3_5  bool   `schema:"lesson_enhmerwsi_ari8mos_spoudaston_pou_aksiologisan_to_ma8ima"`

	// Διδακτέα Ύλη
	Field4_1  string `schema:"lesson_didaktea_yli_anaprosarmogi_mathimatos"`
	Field4_2  string `schema:"lesson_didaktea_yli_epikalipsi_ylis_me_alla_ma8imata"`

	// Διδακτικά Βοηθήματα
	Field5_1  string `schema:"lesson_didaktika_boi8imata_boi8imata_pou_dianemontai_stous_foithtes"`
	Field5_2  string `schema:"lesson_didaktika_boi8imata_epikairopoihsh_ton_boh8imaton"`
	Field5_3  string `schema:"lesson_didaktika_boi8imata_pososto_didaskomenis_ylis_pou_kaliptete"`
	Field5_4  string `schema:"lesson_didaktika_boi8imata_prostheti_bibliografia"`
	Field5_5  string `schema:"lesson_didaktika_boi8imata_epikairopoihsh_ylis"`
	Field5_6  string `schema:"lesson_didaktika_boi8imata_gnostopoihsh_ylis_stous_foithtes"`
}

