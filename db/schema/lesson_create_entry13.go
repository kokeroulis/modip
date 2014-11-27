package DbSchema

const LessonCreateReportFormEntry13 = `
CREATE TABLE LessonCreateReportFormEntry13 (
	lesson int references lesson(id) on delete cascade,
	statistoika_stoixeia_koinopoieitai_katologos_ton_spoudaston_pou_einai_grammenoi_sto_ma8ima text default ' '
)
`
