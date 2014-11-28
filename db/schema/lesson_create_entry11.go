package DbSchema

const LessonCreateReportFormEntry11 = `
CREATE TABLE LessonCreateReportFormEntry11 (
    lesson                                                                                      int references lesson(id) on delete cascade,
	TPE_sth_didaskalia_tou_ma8imatos                text default ' ',
	ma8isiaka_boh8imata_basismena_se_TPE            text default ' ',
	TPE_sthn_ergasthriaki_ekpedeusi                 text default ' ',
	TPE_sthn_aksiologisi_ton_spoudaston_pws         text default ' ',
	TPT_sthn_epikoinonia_sas_me_tous_spoudastes_pws text default ' '
)
`
