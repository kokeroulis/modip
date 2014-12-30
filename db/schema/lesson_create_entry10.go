package DbSchema

const LessonCreateReportFormEntry10 = `
CREATE TABLE LessonCreateReportFormEntry10 (
        lesson int references lesson(id) on delete cascade,
        akademic_year int references akademic_year(id) on delete cascade,
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ai8ouses_didaskalias_pou_xrisimopoiountai text default ' ',
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_ergasthria_pou_xrhsimopoiountai_gia_to_sugkekrimeno_ma8ima text default ' ',
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_einai_dia8esima_ta_ergasthria_tou_ma8imatos_gia_xrisi_ektos_programmatismenon_wrwn text default ' ',
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_spoudasthria text default ' ',
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_xrhsimopoihtai_ekpedeutiko_logismiko_kai_poio text default ' ',
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_yparxei_ikanopoihtiki_ypostiriksi_tou_ma8imatos_apo_tin_biblio8iki text default ' ',
        dia8esimi_ekpedeutikh_ypodomh_tou_ma8imatos_pws_krinete_sunolika_tin_dia8esimi_ekpedeutiki_ypodomi text default ' ',
)
`

