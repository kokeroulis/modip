package DbSchema

const TeacherCreateReportFormEntry2 = `
CREATE TABLE TeacherCreateReportFormEntry2 (
	teacher int references teacher(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	anagnorish_tou_episthmonikou_kai_allou_ergou_eteroanafores                                                int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_anafores_tou_eidikou_episthmonikou_typou                     int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_epitropes_episthmonikon_sunedrion              int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_suntaktikes_epitropes_episthmonikon_sunedrion  int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_kritis_se_episthmonika_periodika                             int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_proskliseis_gia_dialekseis_se_die8nh_sunedria                int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_euresitexnias                                      int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_brabeia                                            int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_timitikoi_titloi                                   int default 0,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_epeksigiseis                                     text default ' '
)
`
