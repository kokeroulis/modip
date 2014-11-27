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


=============== Start create update func
func (f *TeacherCreateReportFormEntry2) Update(lessonId int) {
query := `UPDATE TeacherCreateReportFormEntry2 SET
	anagnorish_tou_episthmonikou_kai_allou_ergou_eteroanafores = $1,
	anagnorish_tou_episthmonikou_kai_allou_ergou_anafores_tou_eidikou_episthmonikou_typou = $2,
	anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_epitropes_episthmonikon_sunedrion = $3,
	anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_suntaktikes_epitropes_episthmonikon_sunedrion = $4,
	anagnorish_tou_episthmonikou_kai_allou_ergou_kritis_se_episthmonika_periodika = $5,
	anagnorish_tou_episthmonikou_kai_allou_ergou_proskliseis_gia_dialekseis_se_die8nh_sunedria = $6,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_euresitexnias = $7,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_brabeia = $8,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_timitikoi_titloi = $9,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_epeksigiseis = $10
WHERE lesson = $`

_, err := Db.Database.Exec(query,
	Field1,
	Field2,
	Field3,
	Field4,
	Field5,
	Field6,
	Field7,
	Field8,
	Field9,
	Field10
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *TeacherCreateReportFormEntry2) Load(lessonId int) {
query := `SELECT
	anagnorish_tou_episthmonikou_kai_allou_ergou_eteroanafores,
	anagnorish_tou_episthmonikou_kai_allou_ergou_anafores_tou_eidikou_episthmonikou_typou,
	anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_epitropes_episthmonikon_sunedrion,
	anagnorish_tou_episthmonikou_kai_allou_ergou_summetoxes_se_suntaktikes_epitropes_episthmonikon_sunedrion,
	anagnorish_tou_episthmonikou_kai_allou_ergou_kritis_se_episthmonika_periodika,
	anagnorish_tou_episthmonikou_kai_allou_ergou_proskliseis_gia_dialekseis_se_die8nh_sunedria,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_euresitexnias,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_brabeia,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_timitikoi_titloi,
	anagnorish_tou_episthmonikou_kai_allou_ergou_diplomata_epeksigiseis
FROM TeacherCreateReportFormEntry2
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5,
		&f.Field6,
		&f.Field7,
		&f.Field8,
		&f.Field9,
		&f.Field10
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
