package DbSchema

const LessonCreateReportFormEntry6 = `
CREATE TABLE LessonCreateReportFormEntry6 (
        lesson int references lesson(id) on delete cascade,
        lesson_alles_ekpedeutikes_drastiriotites_alles_drastiriotites text default ' '
)
`

