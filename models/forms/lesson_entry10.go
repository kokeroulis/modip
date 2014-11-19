package forms

import (
	_"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry10 struct {
	// Διαθέσιμη εκπαιδευτική υποδομή του μαθήματος
	Field1 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ai8ouses_didaskalias_pou_xrisimopoiountai"`
	Field2 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ergasthria_pou_xrhsimopoiountai_gia_to_sugkekrimeno_ma8ima"`
	Field3 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_einai_dia8esima_ta_ergasthria_tou_ma8imatos_gia_xrisi_ektos_programmatismenon_wrwn"`
	Field4 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_spoudasthria"`
	Field5 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_xrhsimopoihtai_ekpedeutiko_logismiko_kai_poio"`
	Field6 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_yparxei_ikanopoihtiki_ypostiriksi_tou_ma8imatos_apo_tin_biblio8iki"`
	Field7 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_pws_krinete_sunolika_tin_dia8esimi_ekpedeutiki_ypodomi"`
	Field8 string `schema:"dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_"`
}

=============== Start create table
CREATE TABLE LessonCreateReportFormEntry10 (
	lesson int references lesson(id) on delete cascade,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ai8ouses_didaskalias_pou_xrisimopoiountai text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ergasthria_pou_xrhsimopoiountai_gia_to_sugkekrimeno_ma8ima text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_einai_dia8esima_ta_ergasthria_tou_ma8imatos_gia_xrisi_ektos_programmatismenon_wrwn text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_spoudasthria text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_xrhsimopoihtai_ekpedeutiko_logismiko_kai_poio text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_yparxei_ikanopoihtiki_ypostiriksi_tou_ma8imatos_apo_tin_biblio8iki text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_pws_krinete_sunolika_tin_dia8esimi_ekpedeutiki_ypodomi text,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ text
)
=============== End Create table
=============== Start create update func
func (f *LessonCreateReportFormEntry10) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry10 SET 
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ai8ouses_didaskalias_pou_xrisimopoiountai = $1,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ergasthria_pou_xrhsimopoiountai_gia_to_sugkekrimeno_ma8ima = $2,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_einai_dia8esima_ta_ergasthria_tou_ma8imatos_gia_xrisi_ektos_programmatismenon_wrwn = $3,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_spoudasthria = $4,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_xrhsimopoihtai_ekpedeutiko_logismiko_kai_poio = $5,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_yparxei_ikanopoihtiki_ypostiriksi_tou_ma8imatos_apo_tin_biblio8iki = $6,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_pws_krinete_sunolika_tin_dia8esimi_ekpedeutiki_ypodomi = $7,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ = $8
WHERE lesson = $`

_, err := Db.Database.Exec(query, 
	Field1, 
	Field2, 
	Field3, 
	Field4, 
	Field5, 
	Field6, 
	Field7, 
	Field8
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *LessonCreateReportFormEntry10) Load(lessonId int) {
query := `SELECT 
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ai8ouses_didaskalias_pou_xrisimopoiountai,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ergasthria_pou_xrhsimopoiountai_gia_to_sugkekrimeno_ma8ima,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_einai_dia8esima_ta_ergasthria_tou_ma8imatos_gia_xrisi_ektos_programmatismenon_wrwn,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_spoudasthria,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_xrhsimopoihtai_ekpedeutiko_logismiko_kai_poio,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_yparxei_ikanopoihtiki_ypostiriksi_tou_ma8imatos_apo_tin_biblio8iki,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_pws_krinete_sunolika_tin_dia8esimi_ekpedeutiki_ypodomi,
	dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_
FROM LessonCreateReportFormEntry10
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO). 
	Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5,
		&f.Field6,
		&f.Field7,
		&f.Field8
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
