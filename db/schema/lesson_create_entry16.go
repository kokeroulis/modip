package DbSchema

const LessonCreateReportFormEntry16 = `
CREATE TABLE LessonCreateReportFormEntry16 (
	lesson int references lesson(id) on delete cascade,
	sxolia_sxolia                     text default ' '
)
`
