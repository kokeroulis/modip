package DbSchema

const LessonCreateReportFormEntry4 = `
CREATE TABLE LessonCreateReportFormEntry4 (
	lesson                                               int references lesson(id) on delete cascade,
	lesson_didaktea_yli_anaprosarmogi_mathimatos         text default ' ',
	lesson_didaktea_yli_epikalipsi_ylis_me_alla_ma8imata text default ' '
)
`
