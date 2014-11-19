package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry2 struct {
	// Αναγνώριση επιστημονικού και άλλου έργου
	Field1 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_eteroanafores"`
    Field2 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_anafores_tou_eidikou_episthmonikou_typou"`
    Field3 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_epitropes_episthmonikon_sunedrion"`
    Field4 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_suntaktikes_epitropes_episthmonikon_sunedrion"`
    Field5 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_kritis_se_episthmonika_periodika"`
    Field6 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_proskliseis_gia_dialekseis_se_die8nh_sunedria"`
    Field7 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_euresitexnias"`
    Field8 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_brabeia"`
    Field9 int      `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_timitikoi_titloi"`
    Field10 string  `schema:"anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_epeksigiseis"`
}

