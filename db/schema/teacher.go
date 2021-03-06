package DbSchema

const Teacher = `
CREATE SEQUENCE seq_teacherIds;

CREATE TABLE teacher
(
    id         int  primary key default nextval('seq_teacherIds'),
    name       text not null,
    email      text not null,
    password   text not null,
    department int  not null references department(id) on delete cascade,
    username   text not null,
    type       int  not null
);
`
