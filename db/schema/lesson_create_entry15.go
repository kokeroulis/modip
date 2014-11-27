package DbSchema

const LessonCreateReportFormEntry15 = `
CREATE TABLE LessonCreateReportFormEntry15 (
	lesson                                             int references lesson(id) on delete cascade,
	apopsi_spoudaston_gia_to_ma8ima_yparxei_diadikasia_aksiologisis_tou_ma8imatos text default ' ',
	apopsi_spoudaston_gia_to_ma8ima_pws_aksiopoiountai_auta_ta_apotelesmata       text default ' '
)
`
