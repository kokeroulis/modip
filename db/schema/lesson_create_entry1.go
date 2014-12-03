package DbSchema

const LessonCreateReportFormEntry1 = `
CREATE TABLE LessonCreateReportFormEntry1 (
	lesson int references lesson(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	lesson_perigrafi_periexomeno_ma8imatos text default ' ',
	lesson_perigrafi_ma8isiakoi_stoxoi     text default ' '
)
`
