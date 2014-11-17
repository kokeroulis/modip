package DbSchema

const LessonCreateReportFormEntry1 = `
CREATE TABLE LessonCreateReportFormEntry1 (
	lesson                                 int references lesson(id) on delete cascade,
	lesson_perigrafi_periexomeno_ma8imatos text,
	lesson_perigrafi_ma8isiakoi_stoxoi     text
)
`
