package DbSchema

const TeacherCreateReportFormEntry5 = `
CREATE TABLE TeacherCreateReportFormEntry5 (
	teacher                                         int references teacher(id) on delete cascade,
	sundesi_me_thn_koinonia_sundesi_me_tin_koinonia text default ' '
)
`
