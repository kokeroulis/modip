package DbSchema

const Lesson = `
CREATE SEQUENCE seq_lessonIds;

CREATE TABLE lesson
(
	id            int     primary key default nextval('seq_lessonIds'),
	name          text    not null,
	department    int     not null references department(id) on delete cascade,
	courseCode    text    not null,
	cardisoftCode text    not null,
	isPostDegree  boolean
);
`
