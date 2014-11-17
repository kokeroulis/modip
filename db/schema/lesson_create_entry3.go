package DbSchema

const LessonCreateReportFormEntry3 = `
CREATE TABLE LessonCreateReportFormEntry3 (
	lesson                                                         int references lesson(id) on delete cascade,
	lesson_enhmerwsi_selida_odigou_spoudon                         int,
	lesson_enhmerwsi_website                                       text,
	lesson_enhmerwsi_aksiologisi_E                                 boolean,
	lesson_enhmerwsi_aksiologisi_8                                 boolean,
	lesson_enhmerwsi_ari8mos_spoudaston_pou_aksiologisan_to_ma8ima boolean
)
`
