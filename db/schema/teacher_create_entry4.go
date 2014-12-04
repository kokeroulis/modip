package DbSchema

const TeacherCreateReportFormEntry4 =`
CREATE SEQUENCE seq_TeacherCreateReportFormEntry4Ids;

CREATE TABLE TeacherCreateReportFormEntry4 (
	id            int primary key default nextval('seq_TeacherCreateReportFormEntry4Ids'),
	teacher       int references teacher(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,

	ereunitikes_ypodomes_arithmos_kai_xwritikotita_ereunitikon_ergasthrion_pou_xrisimopoieitai    text default ' ',
	ereunitikes_ypodomes_eparkeia_katalilotita_kai_poiotita_ton_ereunitikon_ergasthrion           text default ' ',
	ereunitikes_ypodomes_eparkeia_katallilotita_kai_poithta_tou_ergasthriakou_eksoplismou         text default ' ',
	ereunitikes_ypodomes_kalyptoun_oi_dia8esimes_ypodomes_tis_anagkes_tis_ereunhtikis_diadikasias text default ' ',
	ereunitikes_ypodomes_poia_apo_ta_ereunitika_sas_antikeimena_den_kalyptontai_apo_tis_ypodomes  text default ' ',
	ereunitikes_ypodomes_poso_entatiki_xrish_kanete_ton_sugkrekrimenon_ereunitikon_ypodomon       text default ' ',
	ereunitikes_ypodomes_ananeosi_ereunitikon_ypodomon                                            text default ' ',
	ereunitikes_ypodomes_pws_epidiokete_th_xrimatodothsh_gia_promi8eia                            text default ' ',
	ereunitikes_ypodomes_praktiki_akiopoihsh_ton_ereunitikon_apotelesmaton                        text default ' '
)
`
const TeacherCreateReportFormEntry4Helper =`
CREATE TABLE TeacherCreateReportFormEntry4Helper (
	form    int references TeacherCreateReportFormEntry4(id) on delete cascade,
	content text default ' '
)
`
