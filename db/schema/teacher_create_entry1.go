package DbSchema

const TeacherCreateReportFormEntry1 = `
CREATE SEQUENCE seq_TeacherCreateReportFormEntry1Ids;


CREATE TABLE TeacherCreateReportFormEntry1 (
	id               int primary key default nextval('seq_TeacherCreateReportFormEntry1Ids'),
	teacher          int references teacher(id) on delete cascade,
	author           text default ' ',
	title            text default ' ',
	is_magazine      boolean default false,
	publisher        text default ' ',
	publication_date text default ' ',
	type             text default ' '
)
`
