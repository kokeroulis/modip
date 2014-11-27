package DbSchema

const TeacherCreateReportFormEntry5 = `
CREATE TABLE TeacherCreateReportFormEntry5 (
	lesson               nt references lesson(id) on delete cascade,
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia text default ' '
)
`
