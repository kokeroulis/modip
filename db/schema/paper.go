package DbSchema

const Paper = `
CREATE SEQUENCE seq_paperIds;

CREATE TABLE paper
(
    id      int  primary key default nextval('seq_paperIds'),
    title   text unique,
    teacher int  not null references teacher(id) on delete cascade
);
`
