package DbSchema

const TeacherCreateReportFormEntry1 = `
CREATE TABLE TeacherCreateReportFormEntry1 (
	teacher          int references teacher(id) on delete cascade,
	author           text default ' ',
	title            text default ' ',
	is_magazine      boolean default false,
	publisher        text default ' ',
	publication_date text default ' ',
	type             text default ' '
)
`
