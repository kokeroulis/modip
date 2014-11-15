package DbSchema

const LessonCreateForm =`
CREATE TABLE lesson_create_form
(
	lesson   int not null unique references lesson(id) on delete cascade,

	-- Περιγραφή
	field1_1_lesson_perigrafi_periexomeno_ma8imatos                              text,
	field1_2_lesson_perigrafi_ma8isiakoi_stoxoi                                  text,

	-- Διδασκαλίας
	field2_1_lesson_didaskalia_typos_ma8iomatos1                                 text,
	field2_2_lesson_didaskalia_typos_ma8iomatos2                                 text,
	field2_3_lesson_didaskalia_typos_ma8iomatos3                                 text,
	field2_4_lesson_didaskalia_typos_ma8iomatos4                                 text,
	field2_5_lesson_didaskalia_typos_ma8iomatos5                                 text,

	field2_6_lesson_didaskalia_wres_didaskaliasi_8                               text,
	field2_7_lesson_didaskalia_wres_didaskalias_AP                               text,
	field2_8_lesson_didaskalia_wres_didaskalias_E                                text,
	field2_9_lesson_didaskalia_monadesECTS_8AP                                   text,
	field2_10_lesson_didaskalia_monadesECTS_E                                    text,
	field2_11_lesson_didaskalia_wres_didaskalias_ana_bdomada                     text,
	field2_12_lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino      text,
	field2_13_lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli text,
	field2_14_lesson_didaskalia_pollapli_bibiografia                             boolean,
	field2_15_lesson_didaskalia_diati8ete_ergasia_proodos                        boolean,
	field2_16_lesson_didaskalia_ypoxreotiki_ergasia_proodos                      boolean,

	-- Ενημέρωση Αξιολόγηση
	field3_1_lesson_enhmerwsi_selida_odigou_spoudon                              int,
	field3_2_lesson_enhmerwsi_website                                            text,
	field3_3_lesson_enhmerwsi_aksiologisi_E                                      boolean,
	field3_4_lesson_enhmerwsi_aksiologisi_8                                      boolean,
	field3_5_lesson_enhmerwsi_ari8mos_spoudaston_pou_aksiologisan_to_ma8ima      boolean,

	-- Διδακτέα Ύλη
	field4_1_lesson_didaktea_yli_anaprosarmogi_mathimatos                        text,
	field4_2_lesson_didaktea_yli_epikalipsi_ylis_me_alla_ma8imata                text,

	-- Διδακτικά Βοηθήματα
	field5_1_lesson_didaktika_boi8imata_boi8imata_pou_dianemontai_stous_foithtes text,
	field5_2_lesson_didaktika_boi8imata_epikairopoihsh_ton_boh8imaton            text,
	field5_3_lesson_didaktika_boi8imata_pososto_didaskomenis_ylis_pou_kaliptete  text,
	field5_4_lesson_didaktika_boi8imata_prostheti_bibliografia                   text,
	field5_5_lesson_didaktika_boi8imata_epikairopoihsh_ylis                      text,
	field5_6_lesson_didaktika_boi8imata_gnostopoihsh_ylis_stous_foithtes         text
);
`
