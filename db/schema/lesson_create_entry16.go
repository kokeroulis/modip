package DbSchema

const LessonCreateReportFormEntry16 = `
CREATE TABLE LessonCreateReportFormEntry16 (
	lesson int references lesson(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	sxolia_sxolia                     text default ' '
)
`
