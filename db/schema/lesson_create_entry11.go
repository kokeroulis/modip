package DbSchema

const LessonCreateReportFormEntryi11 = `
CREATE TABLE LessonCreateReportFormEntry11 (
    lesson                                                                                      int references lesson(id) on delete cascade,
	aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPE_sth_didaskalia_tou_ma8imatos                text default ' ',
	aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_ma8isiaka_boh8imata_basismena_se_TPE            text default ' ',
	aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPE_sthn_ergasthriaki_ekpedeusi                 text default ' ',
	aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPE_sthn_aksiologisi_ton_spoudaston_pws         text default ' ',
	aksiopoihsh_texnologion_pliroforikhs_kai_epikoinonion_xrhsimopoiountai_TPT_sthn_epikoinonia_sas_me_tous_spoudastes_pws text default ' '
)
`
